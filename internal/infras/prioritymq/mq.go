package prioritymq

import (
	"context"
)

// MQ 优先级队列接口
type MQ interface {
	// Publish 发送消息
	Publish(ctx context.Context, topic string, msg []byte, opts ...PubOption) error

	// Subscribe 消费消息
	Subscribe(ctx context.Context, topics []string, channel string, subHandler SubHandler, opts ...SubOption) error

	// Shutdown 平滑退出
	Shutdown(ctx context.Context) error
}

// SubHandler subscribe func
type SubHandler func(ctx context.Context, msg []byte) error
