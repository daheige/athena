package monitor

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/daheige/athena/internal/infras/monitor/gpprof"
)

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,world"))
}

// http://localhost:2339/debug/pprof/
// http://localhost:1339/hello
// http://localhost:1339/test
// http://localhost:2339/metrics

// TestMonitorHandler test monitor.
func TestMonitorHandler(t *testing.T) {
	// 添加prometheus性能监控指标
	prometheus.MustRegister(WebRequestTotal)
	prometheus.MustRegister(WebRequestDuration)

	prometheus.MustRegister(CpuTemp)
	prometheus.MustRegister(HdFailures)

	// 性能监控的端口port+1000,只能在内网访问
	httpMux := gpprof.New()

	port := 1339
	// 添加prometheus metrics处理器
	httpMux.Handle("/metrics", promhttp.Handler())
	gpprof.Run(httpMux, uint16(port+1000))

	router := http.NewServeMux()
	router.HandleFunc("/test", MonitorHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("ok"))
	}))

	router.Handle("/hello", MonitorHandler(&helloHandler{}))

	// 服务server设置
	server := &http.Server{
		Handler:           router,
		Addr:              fmt.Sprintf("0.0.0.0:%d", port),
		IdleTimeout:       20 * time.Second, // tcp idle time
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	// 在独立携程中运行
	log.Println("server run on: ", port)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// mock server shutdown.
	time.Sleep(200 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	<-ctx.Done()

	log.Println("server shutdown")
}
