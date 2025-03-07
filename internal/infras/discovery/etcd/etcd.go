// Package etcd impl registry interface.
package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/daheige/athena/internal/infras/discovery"
)

var _ discovery.Registry = (*etcdImpl)(nil)

type etcdImpl struct {
	endpoints   []string      // etcd节点列表
	dialTimeout time.Duration // 默认5s
	client      *clientv3.Client

	meta              *registerMeta
	prefix            string // 默认为/services
	keepaliveInterval time.Duration

	// etcd 用户名和密码可选
	username string
	password string
}

type registerMeta struct {
	leaseID clientv3.LeaseID
	stop    chan struct{}
}

// Option etcdImpl functional option
type Option func(*etcdImpl)

// WithDialTimeout 设置 dialTimeout
func WithDialTimeout(dialTimeout time.Duration) Option {
	return func(impl *etcdImpl) {
		impl.dialTimeout = dialTimeout
	}
}

// WithUsername 设置 username
func WithUsername(username string) Option {
	return func(impl *etcdImpl) {
		impl.username = username
	}
}

// WithPassword 设置 password
func WithPassword(password string) Option {
	return func(impl *etcdImpl) {
		impl.password = password
	}
}

// WithPrefix 设置服务前缀
func WithPrefix(prefix string) Option {
	return func(impl *etcdImpl) {
		impl.prefix = prefix
	}
}

// WithKeepaliveInterval 设置keepalive时间间隔
func WithKeepaliveInterval(interval time.Duration) Option {
	return func(impl *etcdImpl) {
		impl.keepaliveInterval = interval
	}
}

// New 创建一个服务注册和发现的实例
func New(endpoints []string, opts ...Option) (discovery.Registry, error) {
	impl := &etcdImpl{
		endpoints:         endpoints,
		dialTimeout:       5 * time.Second,
		prefix:            "athena/registry-etcd",
		keepaliveInterval: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(impl)
	}

	config := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: impl.dialTimeout,
		Username:    impl.username,
		Password:    impl.password,
	}
	client, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	impl.client = client
	return impl, nil
}

// Register 注册服务
func (e *etcdImpl) Register(s discovery.Service, ttl ...time.Duration) error {
	if s.InstanceID == "" {
		s.InstanceID = uuid.New().String()
	}
	if s.Created == "" {
		s.Created = time.Now().Format("2006-01-02 15:04:05")
	}
	if s.Address == "" {
		return errors.New("address invalid,please provide service address，eg: ip:port")
	}

	var ttlTime int64 = 20
	if len(ttl) > 0 && ttl[0] > 0 {
		ttlTime = int64(ttl[0].Seconds())
	}

	leaseID, err := e.grantLease(ttlTime)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%s/%s/%s", e.prefix, s.Name, s.InstanceID)
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err = e.client.Put(ctx, key, string(b), clientv3.WithLease(leaseID))
	if err != nil {
		return err
	}

	e.meta = &registerMeta{
		leaseID: leaseID,
		stop:    make(chan struct{}, 1),
	}
	err = e.keepalive(e.meta)
	if err != nil {
		return err
	}

	log.Printf("register service:%v leaseID:%v instaceID:%v success\n", s.Name, leaseID, s.InstanceID)
	go func() {
		ticker := time.NewTicker(e.keepaliveInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err2 := e.keepalive(e.meta)
				if err2 != nil {
					log.Printf(
						"keep alive service:%v leaseID:%v instaceID:%v failed,error: %v\n",
						s.Name, leaseID, s.InstanceID, err2,
					)
				} else {
					log.Printf(
						"keep alive service:%v leaseID:%v instaceID:%v success\n",
						s.Name, leaseID, s.InstanceID,
					)
				}
			case <-e.meta.stop:
				log.Printf(
					"service:%v leaseID:%v instaceID:%v has been deregistered\n",
					s.Name, leaseID, s.InstanceID,
				)
				return
			default:
			}
		}
	}()

	return nil
}

func (e *etcdImpl) keepalive(meta *registerMeta) error {
	keepAlive, err := e.client.KeepAlive(context.Background(), meta.leaseID)
	if err != nil {
		return err
	}

	go func() {
		// eat keepAlive channel to keep related lease alive.
		for range keepAlive {
			select {
			case <-meta.stop:
				log.Printf("stop keepalive lease %v for etcd registry\n", meta.leaseID)
				return
			default:
			}
		}
	}()

	return nil
}

// 创建租约
func (e *etcdImpl) grantLease(ttl int64) (clientv3.LeaseID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := e.client.Grant(ctx, ttl)
	if err != nil {
		return clientv3.NoLease, err
	}

	return resp.ID, nil
}

// Deregister 移除服务
func (e *etcdImpl) Deregister(name string, instanceID string) error {
	key := fmt.Sprintf("%s/%s/%s", e.prefix, name, instanceID)
	_, err := e.client.Delete(context.Background(), key)
	if err != nil {
		return err
	}

	log.Printf("deregister service:%v instaceID:%v success\n", name, instanceID)
	if e.meta != nil {
		close(e.meta.stop)
	}

	return nil
}

// GetServices 通过name获取服务列表
func (e *etcdImpl) GetServices(name string) ([]*discovery.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	key := fmt.Sprintf("%s/%s", e.prefix, name)
	resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, discovery.ErrServiceNotFound
	}

	// 获取所有的服务实例列表
	services := make([]*discovery.Service, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		serviceEntry := &discovery.Service{}
		err = json.Unmarshal(kv.Value, serviceEntry)
		if err != nil {
			log.Printf("unmarshal service failed,error:%v", err)
			continue
		}

		services = append(services, serviceEntry)
	}

	return services, nil
}

// String 实现服务注册和发现的名字
func (e *etcdImpl) String() string {
	return "etcd"
}
