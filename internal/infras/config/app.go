package config

import (
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/daheige/athena/internal/infras/setting"
)

// AppConfig 应用基本配置
type AppConfig struct {
	// 这里的 mapstructure tag标签用于指定配置文件的字段名字
	AppDebug        bool          `mapstructure:"app_debug"`         // 是否开启调试模式
	AppEnv          string        `mapstructure:"app_env"`           //  prod,test,local,dev
	HttpPort        uint16        `mapstructure:"http_port"`         // http 服务端口
	MonitorPort     uint16        `mapstructure:"monitor_port"`      // http 服务端口
	GrpcPort        uint16        `mapstructure:"grpc_port"`         // grpc 服务端口
	GrpcGatewayPort uint16        `mapstructure:"grpc_gateway_port"` // grpc http gateway 服务端口
	GracefulWait    time.Duration `mapstructure:"graceful_wait"`     // 平滑退出等待时间

	// 数据库配置
	DbConf DbConf `mapstructure:"db_conf"`

	// redis配置
	RedisConf RedisConf `mapstructure:"redis_conf"`

	// 其他情况根据实际情况添加
}

// 配置文件读取的接口
var conf setting.Config

// InitAppConfig 初始化app config
// 这个函数的代码可以根据实际情况在main.go初始化
func InitAppConfig() *AppConfig {
	var err error
	// 读取配置文件，并初始化redis和mysql
	conf, err = Load("./app.yaml")
	if err != nil {
		log.Fatalln("failed to load config:", err)
	}

	appConfig := &AppConfig{}
	err = conf.ReadSection("app_conf", appConfig)
	// log.Println("read app_conf err: ", err)
	if appConfig.AppDebug {
		log.Println("app_conf:", appConfig)
	}

	return appConfig
}

// InitDB 根据配置文件配置的名字获取DB句柄
func InitDB(name string) *gorm.DB {
	dbConfig := DbConf{}
	err := conf.ReadSection(name, &dbConfig)
	// log.Println("db conf:", dbConfig)
	if err != nil {
		log.Fatalln("failed to load db config:", err)
	}

	// 建立mysql连接
	db, err := dbConfig.InitDB()
	if err != nil {
		log.Fatalln("failed to init db:", err)
	}

	return db
}

// InitRedisClient 获取redis client
func InitRedisClient(name string) redis.UniversalClient {
	redisConfig := RedisConf{}
	err := conf.ReadSection(name, &redisConfig)
	// log.Println("redis conf:", redisConfig)
	if err != nil {
		log.Fatalln("failed to load redis config:", err)
	}

	return redisConfig.InitClient()
}
