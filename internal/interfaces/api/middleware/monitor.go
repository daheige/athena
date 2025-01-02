package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/daheige/athena/internal/infras/monitor"
)

// WrapMonitor metrics 性能监控
// gin处理器函数，包装 handler function,不侵入业务逻辑
func WrapMonitor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		// counter类型 metrics 的记录方式
		monitor.WebRequestTotal.With(prometheus.Labels{
			"method": ctx.Request.Method, "endpoint": ctx.Request.URL.Path,
		}).Inc()

		// Histogram类型 metrics 的记录方式
		monitor.WebRequestDuration.With(prometheus.Labels{
			"method": ctx.Request.Method, "endpoint": ctx.Request.URL.Path,
		}).Observe(time.Since(start).Seconds())
	}
}
