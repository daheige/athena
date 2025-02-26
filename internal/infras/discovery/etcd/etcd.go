// Package etcd impl registry interface.
package etcd

import (
	"context"
	"encoding/json"
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

	meta   *registerMeta
	prefix string // 默认为/services

	// etcd 用户名和密码可选
	username string
	password string
}

type registerMeta struct {
	leaseID clientv3.LeaseID
	ctx     context.Context
	cancel  context.CancelFunc
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

// New 创建一个服务注册和发现的实例
func New(endpoints []string, opts ...Option) (discovery.Registry, error) {
	impl := &etcdImpl{
		endpoints:   endpoints,
		dialTimeout: 5 * time.Second,
		prefix:      "/services",
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

	var ttlTime int64 = 10
	if len(ttl) > 0 && ttl[0] > 0 {
		ttlTime = int64(ttl[0].Seconds())
	}

	leaseID, err := e.grantLease(ttlTime)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	key := fmt.Sprintf("%s/%s/%s", e.prefix, s.Name, s.InstanceID)
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}

	// log.Printf("register serviceName:%v leaseID:%v instaceID:%v\n", s.Name, leaseID, s.InstanceID)
	_, err = e.client.Put(ctx, key, string(b), clientv3.WithLease(leaseID))
	if err != nil {
		return err
	}

	meta := &registerMeta{
		leaseID: leaseID,
	}
	meta.ctx, meta.cancel = context.WithCancel(context.Background())
	err = e.keepalive(meta)
	if err != nil {
		return err
	}

	e.meta = meta

	log.Printf("register serviceName:%v leaseID:%v instaceID:%v success\n", s.Name, leaseID, s.InstanceID)
	return nil
}

func (e *etcdImpl) keepalive(meta *registerMeta) error {
	keepAlive, err := e.client.KeepAlive(meta.ctx, meta.leaseID)
	if err != nil {
		return err
	}

	go func() {
		// eat keepAlive channel to keep related lease alive.
		log.Printf("start keepalive lease %x for etcd registry\n", meta.leaseID)
		for range keepAlive {
			select {
			case <-meta.ctx.Done():
				log.Printf("stop keepalive lease %x for etcd registry\n", meta.leaseID)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	key := fmt.Sprintf("%s/%s/%s", e.prefix, name, instanceID)
	log.Printf("deregister serviceName:%v instaceID:%v key:%s\n", name, instanceID, key)
	_, err := e.client.Delete(ctx, key)
	if err != nil {
		return err
	}

	if e.meta != nil {
		e.meta.cancel()
	}

	log.Printf("deregister serviceName:%v instaceID:%v success\n", name, instanceID)
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
