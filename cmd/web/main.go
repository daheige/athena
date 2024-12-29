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

	"github.com/daheige/athena/internal/infras/config"
	"github.com/daheige/athena/internal/infras/logger"
	"github.com/daheige/athena/internal/infras/monitor"
	"github.com/daheige/athena/internal/interfaces/api/routers"
	"github.com/daheige/athena/internal/providers"
)

func main() {
	log.Println("athena web...")
	// 读取配置文件，并初始化redis和mysql
	conf := config.InitAppConfig()

	// 初始化日志配置，如果app_debug = true 将日志输出到终端中
	logger.Default(logger.WithLogFilename("athena-web.log"), logger.WithStdout(conf.AppDebug))

	// 初始化db和redis
	gormClient := config.InitDB("db_conf")
	redisClient := config.InitRedisClient("redis_conf")

	// 初始化路由规则
	router := gin.New()

	// 初始化repos
	repos := providers.NewRepositories(gormClient, redisClient)

	// 注册路由规则
	routers.InitRouters(router, repos)

	// 服务server设置
	server := &http.Server{
		Handler:           router,
		Addr:              fmt.Sprintf("0.0.0.0:%d", conf.HttpPort),
		IdleTimeout:       20 * time.Second, // tcp idle time
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	// 在独立携程中运行
	log.Println("server run on: ", conf.HttpPort)
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

	// 初始化prometheus和pprof
	// 访问地址：http://localhost:2337/metrics
	// 访问地址：http://localhost:2337/debug/pprof/
	monitor.InitMonitor(conf.MonitorPort, true)

	// server平滑重启
	ch := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// receive signal to exit main goroutine.
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	sig := <-ch

	log.Println("exit signal: ", sig.String())
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), conf.GracefulWait)
	defer cancel()

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
