package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	gMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DbConf 数据库配置结构体
type DbConf struct {
	Host      string // 对应ip地址
	Port      int    // 默认3306
	User      string
	Password  string
	Database  string
	Charset   string // 字符集 utf8mb4 支持表情符号
	Collation string // 整理字符集 utf8mb4_unicode_ci

	UsePool      bool // 当前db实例是否采用db连接池,默认不采用，如采用请求配置该参数
	MaxIdleConns int  // 空闲pool个数
	MaxOpenConns int  // 最大open connection个数

	// sets the maximum amount of time a connection may be reused.
	// 设置连接可以重用的最大时间
	// 给db设置一个超时时间，时间小于数据库的超时时间
	MaxLifetime time.Duration // 数据库超时时间

	// 连接超时/读取超时/写入超时设置
	Timeout      time.Duration // Dial timeout
	ReadTimeout  time.Duration // I/O read timeout
	WriteTimeout time.Duration // I/O write timeout

	ParseTime bool   // 格式化时间类型
	Loc       string // 时区字符串 Local,PRC

	ShowSql bool // sql语句是否输出

	// sql输出logger句柄接口
	// logger.Writer 接口需要实现Printf(string, ...interface{}) 方法
	// 具体可以看gorm v2 logger包源码
	// https://github.com/go-gorm/gorm
	Logger logger.Writer `mapstructure:"-"`

	// gorm v2版本新增参数
	gMysqlConfig gMysql.Config `mapstructure:"-"` // gorm v2新增参数gMysql.Config
	gormConfig   gorm.Config   `mapstructure:"-"` // gorm v2新增参数gorm.Config
	LoggerConfig logger.Config `mapstructure:"-"` // gorm v2新增参数logger.Config
}

// InitDB 初始化DB实例对象
func (c *DbConf) InitDB() (*gorm.DB, error) {
	// 是否输出sql日志
	// 这里重写了之前的gorm v1版本的日志输出模式
	log.Println("show sql: ", c.ShowSql)
	if c.ShowSql {
		// 日志对象接口
		var dbLogger logger.Interface
		if c.Logger == nil {
			dbLogger = logger.Default.LogMode(logger.Info)
		} else {
			dbLogger = logger.New(c.Logger, c.LoggerConfig)
		}

		// 设置gorm logger句柄对象
		c.gormConfig.Logger = dbLogger
	} else {
		c.gormConfig.Logger = logger.Discard // 默认是不输出sql
	}

	var err error
	if c.gMysqlConfig.DSN == "" {
		dsn, err := c.dsn()
		if err != nil {
			log.Println("mysql dsn format error: ", err)
			return nil, err
		}

		c.gMysqlConfig.DSN = dsn
	}

	// 下面这种方式实例的gorm.DB 很多参数都没法正确设置，不推荐这么实例化
	// db, err := gorm.Open(gMysql.Open(c.gMysqlConfig.DSN), &gorm.Config{
	// 	Logger: c.gormConfig.Logger,
	// })

	// 对于golang的官方sql引擎，sql.open并非立即连接db,用的时候才会真正的建立连接
	// 但是gorm.Open在设置完db对象后，还发送了一个Ping操作，判断连接是否连接上去
	// 具体可以看gorm/main.go源码Open方法
	db, err := gorm.Open(gMysql.New(c.gMysqlConfig), &c.gormConfig)
	if err != nil {
		log.Println("open mysql connection error: ", err)
		return nil, err
	}

	// 设置连接池
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		log.Println("get sql db error: ", err)
		return nil, err
	}

	if c.UsePool {
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	}

	// 设置连接可以重用的最大存活时间，时间小于数据库的超时时间
	if c.MaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(c.MaxLifetime)
	}

	return db, nil
}

// 生成mysql dsn句柄
func (c *DbConf) dsn() (string, error) {
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}

	if c.Port == 0 {
		c.Port = 3306
	}

	if c.Charset == "" {
		c.Charset = "utf8mb4"
	}

	// 默认字符序，定义了字符的比较规则
	if c.Collation == "" {
		c.Collation = "utf8mb4_general_ci"
	}

	if c.Loc == "" {
		c.Loc = "Local"
	}

	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}

	if c.WriteTimeout == 0 {
		c.WriteTimeout = 5 * time.Second
	}

	if c.ReadTimeout == 0 {
		c.ReadTimeout = 5 * time.Second
	}

	// mysql connection time loc.
	loc, err := time.LoadLocation(c.Loc)
	if err != nil {
		return "", err
	}

	// mysql config
	mysqlConf := mysql.Config{
		User:   c.User,
		Passwd: c.Password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		DBName: c.Database,
		// Connection parameters
		Params: map[string]string{
			"charset": c.Charset,
		},
		Collation:            c.Collation,
		Loc:                  loc,            // Location for time.Time values
		Timeout:              c.Timeout,      // Dial timeout
		ReadTimeout:          c.ReadTimeout,  // I/O read timeout
		WriteTimeout:         c.WriteTimeout, // I/O write timeout
		AllowNativePasswords: true,           // Allows the native password authentication method
		ParseTime:            c.ParseTime,    // Parse time values to time.Time
	}

	return mysqlConf.FormatDSN(), nil
}
