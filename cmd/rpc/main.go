package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/daheige/gmicro/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/daheige/athena/internal/application"
	"github.com/daheige/athena/internal/infras/config"
	"github.com/daheige/athena/internal/infras/discovery"
	"github.com/daheige/athena/internal/infras/discovery/etcd"
	"github.com/daheige/athena/internal/infras/logger"
	"github.com/daheige/athena/internal/infras/monitor"
	"github.com/daheige/athena/internal/interfaces/rpc"
	"github.com/daheige/athena/internal/interfaces/rpc/interceptor"
	"github.com/daheige/athena/internal/pb"
	"github.com/daheige/athena/internal/providers"
)

var (
	shutdownFunc func()
	conf         *config.AppConfig
)

func init() {
	// 读取配置文件，并初始化redis和mysql
	conf = config.InitAppConfig()

	// 初始化日志配置
	logger.Default(logger.WithLogFilename("athena-rpc.log"), logger.WithStdout(conf.AppDebug))

	// 服务退出前的处理函数
	shutdownFunc = func() {
		fmt.Println("Server shutting down")
	}

	// 初始化prometheus和pprof，可以根据实际情况更改
	// monitor.InitMonitor(conf.MonitorPort)
	monitor.InitMonitor(conf.MonitorPort + 1000)
}

func main() {
	// add the /test endpoint
	health := gmicro.Route{
		Method: "GET",
		Path:   "/healthz",
		Handler: func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			now := time.Now().Format("2006-01-02 15:04:05")
			m := map[string]interface{}{
				"code": 0,
				"data": now,
			}

			b, _ := json.Marshal(m)
			_, _ = w.Write(b)
		},
	}

	// 服务注册
	if conf.EnableDiscovery {
		log.Println("conf.Discovery.TargetType: ", conf.Discovery.TargetType)
		r, err := etcd.New(conf.Discovery.Endpoints)
		if err != nil {
			log.Fatal("init service registry error: ", err)
		}
		serviceName := "athena_grpc"
		instanceID := uuid.New().String()
		err = r.Register(discovery.Service{
			Name:       serviceName,
			Version:    "v1",
			InstanceID: instanceID,
			Address:    fmt.Sprintf("localhost:%d", conf.GrpcPort),
		})
		if err != nil {
			log.Fatal("register service error:", err)
		}

		shutdownFunc = func() {
			if e := r.Deregister(serviceName, instanceID); e != nil {
				log.Println("deregister service error:", e)
			}
		}
	}

	opts := []gmicro.Option{
		gmicro.WithRouteOpt(health),
		gmicro.WithShutdownFunc(shutdownFunc),
		gmicro.WithShutdownTimeout(5 * time.Second),
		gmicro.WithHandlerFromEndpoint(pb.RegisterGreeterServiceHandlerFromEndpoint),
		gmicro.WithRequestAccess(true),
		gmicro.WithPrometheus(true),
		gmicro.WithGRPCServerOption(grpc.ConnectionTimeout(10 * time.Second)),
		gmicro.WithGRPCNetwork("tcp"), // grpc server start network

		// 自定义拦截器
		gmicro.WithUnaryInterceptor(interceptor.AccessLog),
	}

	if conf.AppDebug { // 调试模式输出日志到终端中
		opts = append(opts, gmicro.WithLogger(gmicro.LoggerFunc(log.Printf)))
	}

	// micro Option func
	s := gmicro.NewService(opts...)

	// init userService by providers
	gormClient := config.InitDB("db_conf")
	redisClient := config.InitRedisClient("redis_conf")

	// init repos and user service
	repos := providers.NewRepositories(gormClient, redisClient)
	userService := application.NewUserService(repos.UserRepo, repos.UserCache)

	// register grpc service
	pb.RegisterGreeterServiceServer(s.GRPCServer, rpc.NewGreeterService(userService))

	// 这里可以手动添加其他http 路由地址
	newRoute := gmicro.Route{
		Method: "GET",
		Path:   "/info",
		Handler: func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ok"))
		},
	}

	s.AddRoute(newRoute)

	// 服务启动（这里gmicro底层框架已经处理了平滑退出，不需要再处理了）
	// log.Println("Starting http server and grpc server listening on ", conf.GrpcPort)
	// you can start grpc server and http gateway on one port
	log.Fatalln(s.StartGRPCAndHTTPServer(int(conf.GrpcPort)))

	// you can also specify ports for grpc and http gw separately
	// log.Fatalln(s.Start(int(conf.GrpcGatewayPort), int(conf.GrpcPort)))

	// you can start server without http gateway
	// log.Fatalln(s.StartGRPCWithoutGateway(int(conf.GrpcPort)))
}
