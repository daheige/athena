// Package discovery service register and discover interface.
package discovery

import (
	"errors"
	"math/rand"
	"time"
)

// ErrServiceNotFound 服务列表为空
var ErrServiceNotFound = errors.New("services not found")

// Registry 服务发现和注册接口定义
type Registry interface {
	// Register 注册服务
	Register(s Service, ttl ...time.Duration) error

	// Deregister 移除服务
	Deregister(name string, instanceID string) error

	// GetServices 通过name获取服务列表
	GetServices(name string) ([]*Service, error)

	// String 实现服务注册和发现的名字
	String() string
}

// Service 服务基本信息
type Service struct {
	// 服务名字
	Name string `json:"name"`

	// 服务地址，一般来说由host:port组成
	Address string `json:"address"`

	// 服务的唯一标识，例如uuid字符串
	InstanceID string `json:"instance_id"`

	// 当前版本
	Version string `json:"version"`

	// 创建时间
	Created string `json:"created"`

	// 服务的其他元信息
	// Metadata map[string]string{} `json:"metadata"`
	Metadata map[string]interface{} `json:"metadata"`
}

// RoundRobinService 随机挑选一个服务
func RoundRobinService(services []*Service) *Service {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个随机索引
	randomIndex := r.Intn(len(services))
	// 通过随机索引获取元素
	service := services[randomIndex]
	return service
}
