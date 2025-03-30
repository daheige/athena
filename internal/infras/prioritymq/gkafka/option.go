package gkafka

import (
	"time"

	"github.com/daheige/athena/internal/infras/prioritymq"
)

// Options kafka options
type Options struct {
	brokers  []string // client connection address list
	user     string   // user
	password string   // password

	// All three of the below configurations are similar to the
	// `socket.timeout.ms` setting in JVM kafka. All of them default
	// to 30 seconds.
	dialTimeout time.Duration

	// producer timeout default:10s
	producerTimeout time.Duration

	// graceful exit time
	// default:3s
	gracefulWait time.Duration

	// consumer auto commit interval (default: 3s)
	consumerAutoCommitInterval time.Duration

	// The initial offset to use if no offset was previously committed.
	// Should be OffsetNewest = -1 or OffsetOldest = -2. Defaults to OffsetNewest.
	// default:-1
	consumerOffsetsInitial int64

	// logger
	logger prioritymq.Logger
}

// Option functional option for kafka options
type Option func(*Options)

// WithBrokers 设置brokers
func WithBrokers(brokers []string) Option {
	return func(o *Options) {
		o.brokers = brokers
	}
}

// WithUser 设置user
func WithUser(user string) Option {
	return func(o *Options) {
		o.user = user
	}
}

// WithPassword 设置pwd
func WithPassword(password string) Option {
	return func(o *Options) {
		o.password = password
	}
}

// WithDialTimeout 设置dialTimeout
func WithDialTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.dialTimeout = d
	}
}

// WithProducerTimeout 设置producerTimeout
func WithProducerTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.producerTimeout = d
	}
}

// WithLogger 设置logger
func WithLogger(logger prioritymq.Logger) Option {
	return func(o *Options) {
		o.logger = logger
	}
}

// WithGracefulWait 设置gracefulWait
func WithGracefulWait(gracefulWait time.Duration) Option {
	return func(o *Options) {
		o.gracefulWait = gracefulWait
	}
}

// WithConsumerAutoCommitInterval 设置consumerAutoCommitInterval
func WithConsumerAutoCommitInterval(t time.Duration) Option {
	return func(o *Options) {
		o.consumerAutoCommitInterval = t
	}
}

// WithConsumerOffsetsInitial 设置consumerOffsetsInitial
func WithConsumerOffsetsInitial(d int64) Option {
	return func(o *Options) {
		o.consumerOffsetsInitial = d
	}
}
