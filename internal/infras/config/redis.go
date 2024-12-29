package config

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisConf redis配置
type RedisConf struct {
	// host:port address.
	Address string

	// Optional password. Must match the password specified in the
	// requirepass server configuration option.
	Password string

	// Database to be selected after connecting to the server.
	DB int

	// Maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration

	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration

	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration

	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int

	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration

	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int

	// ConnMaxIdleTime is the maximum amount of time a connection may be idle.
	// Should be less than server's timeout.
	//
	// Expired connections may be closed lazily before reuse.
	// If d <= 0, connections are not closed due to a connection's idle time.
	//
	// Default is 30 minutes. -1 disables idle timeout check.
	ConnMaxIdleTime time.Duration

	// ConnMaxLifetime is the maximum amount of time a connection may be reused.
	//
	// Expired connections may be closed lazily before reuse.
	// If <= 0, connections are not closed due to a connection's age.
	//
	// Default is to not close idle connections.
	ConnMaxLifetime time.Duration
}

// InitClient 初始化redis client
func (conf *RedisConf) InitClient() redis.UniversalClient {
	if conf.ConnMaxLifetime == 0 {
		conf.ConnMaxLifetime = 30 * 60 * time.Second
	}

	if conf.ConnMaxIdleTime == 0 {
		conf.ConnMaxIdleTime = 30 * 60 * time.Second
	}

	if conf.DialTimeout == 0 {
		conf.DialTimeout = 5 * time.Second
	}

	if conf.WriteTimeout == 0 {
		conf.WriteTimeout = 3 * time.Second
	}

	if conf.ReadTimeout == 0 {
		conf.ReadTimeout = 3 * time.Second
	}

	opt := &redis.Options{
		Addr:            conf.Address,
		Password:        conf.Password,
		DB:              conf.DB, // use default DB
		MaxRetries:      conf.MaxRetries,
		DialTimeout:     conf.DialTimeout,  // Default is 5 seconds.
		ReadTimeout:     conf.ReadTimeout,  // Default is 3 seconds.
		WriteTimeout:    conf.WriteTimeout, // Default is ReadTimeout.
		PoolSize:        conf.PoolSize,
		PoolTimeout:     conf.PoolTimeout,
		MinIdleConns:    conf.MinIdleConns,
		ConnMaxIdleTime: conf.ConnMaxIdleTime,
		ConnMaxLifetime: conf.ConnMaxLifetime,
	}

	return redis.NewClient(opt)
}
