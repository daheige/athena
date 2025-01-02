package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/daheige/athena/internal/infras/config"
	"github.com/daheige/athena/internal/infras/logger"
	"github.com/daheige/athena/internal/infras/monitor"
	"github.com/daheige/athena/internal/interfaces/api/middleware"
	"github.com/daheige/athena/internal/pb"
)

func main() {
	// 读取配置文件，并初始化redis和mysql
	conf := config.InitAppConfig()

	// 初始化日志配置
	logger.Default(logger.WithLogFilename("athena-gateway.log"), logger.WithStdout(conf.AppDebug))

	// 初始化prometheus和pprof，可以根据实际情况更改
	// monitor.InitMonitor(conf.MonitorPort)
	// 访问地址：http://localhost:9091/metrics
	monitor.InitMonitor(conf.GrpcGatewayPort+1000, true)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// 这个grpc微服务地址，一般来说是一个远程的ip:port，可以根据实际情况更改
	// gRPCAddress := fmt.Sprintf("0.0.0.0:%d",conf.GrpcPort)
	gRPCAddress := fmt.Sprintf("0.0.0.0:%d", conf.GrpcPort)
	err := pb.RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, gRPCAddress, opts)
	if err != nil {
		logger.Fatal(ctx, "failed to register grpc endpoint", map[string]interface{}{
			"trace_error": err.Error(),
		})
	}

	router := gin.Default()
	initRouter(router, mux)

	// 服务server设置
	server := &http.Server{
		Handler:           router,
		Addr:              fmt.Sprintf("0.0.0.0:%d", conf.GrpcGatewayPort),
		IdleTimeout:       20 * time.Second, // tcp idle time
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	// 在独立携程中运行
	log.Println("server run on: ", conf.GrpcGatewayPort)
	go func() {
		defer logger.Recover(context.Background(), "server start panic")

		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Error(context.Background(), "server close error", map[string]interface{}{
					"trace_error": err.Error(),
				})

				log.Println("server close error:", err)
				return
			}

			log.Println("server will exit...")
		}
	}()

	// server平滑重启
	ch := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// receive signal to exit main goroutine.
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	sig := <-ch

	log.Println("exit signal: ", sig.String())
	// Create a deadline to wait for.
	ctx, cancel2 := context.WithTimeout(context.Background(), conf.GracefulWait)
	defer cancel2()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// if your application should wait for other services
	// to finalize based on context cancellation.
	done := make(chan struct{}, 1)
	go func() {
		defer close(done)

		_ = server.Shutdown(ctx)
	}()

	<-done
	<-ctx.Done()

	log.Println("server shutting down")
}

func initRouter(router *gin.Engine, mux *runtime.ServeMux) {
	// 访问日志中间件处理
	logWare := &middleware.LogWare{}

	// 对所有的请求进行性能监控，一般来说生产环境，可以对指定的接口做性能监控
	router.Use(logWare.Access(), logWare.Recover(), middleware.TimeoutHandler(10*time.Second))

	// gin 框架prometheus接入
	router.Use(middleware.WrapMonitor())

	// 路由找不到的情况
	router.NoRoute(middleware.NotFoundHandler())

	// gateway http proxy
	// 这里将proto文件中的路由地址进行路由注册
	// 访问方式：http://localhost:8081/v1/say/1
	router.Any("/v1/*any", gin.WrapH(mux))

	// 自定义路由地址
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
