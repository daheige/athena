package middleware

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/daheige/athena/internal/infras/gutils"
	"github.com/daheige/athena/internal/infras/logger"
)

// LogWare 日志中间件
type LogWare struct{}

// Access 记录访问日志
func (ware *LogWare) Access() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		// uri := ctx.Request.RequestURI

		// 性能分析后发现log.Println输出需要分配大量的内存空间,而且每次写入都需要枷锁处理
		// log.Println("request before")
		// log.Println("request uri: ", uri)

		// 如果采用了nginx x-request-id功能，可以获得x-request-id
		logId := ctx.GetHeader("X-Request-Id")
		if logId == "" {
			logId = gutils.RndUuid() // 日志id
		}

		// 设置跟请求相关的ctx信息
		ctx.Request = gutils.SetValueToHTTPCtx(ctx.Request, logger.XRequestID, logId)
		ctx.Request = gutils.SetValueToHTTPCtx(ctx.Request, logger.ReqClientIP, ctx.ClientIP())
		ctx.Request = gutils.SetValueToHTTPCtx(ctx.Request, logger.RequestURI, ctx.Request.RequestURI)
		ctx.Request = gutils.SetValueToHTTPCtx(ctx.Request, logger.UserAgent, ctx.GetHeader("User-Agent"))
		ctx.Request = gutils.SetValueToHTTPCtx(ctx.Request, logger.RequestMethod, ctx.Request.Method)

		// 记录请求日志
		logger.Info(ctx.Request.Context(), "exec begin", nil)

		ctx.Next()

		// log.Println("request end")
		// 请求结束记录日志
		c := map[string]interface{}{
			"exec_time": time.Now().Sub(t).Seconds(),
		}

		if code := ctx.Writer.Status(); code != 200 {
			c["response_code"] = code
		}

		logger.Info(ctx.Request.Context(), "exec end", c)
	}
}

// Recover gin请求处理中遇到异常或panic捕获
func (ware *LogWare) Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// log.Printf("error:%v", err)
				c := ctx.Request.Context()
				logger.DPanic(c, "exec panic", map[string]interface{}{
					"trace_error": fmt.Sprintf("%v", err),
					"full_stack":  string(gutils.CatchStack()),
				})

				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errMsg := strings.ToLower(se.Error())
						// 记录操作日志
						logger.DPanic(c, "os syscall error", map[string]interface{}{
							"trace_error": errMsg,
						})

						if strings.Contains(errMsg, "broken pipe") ||
							strings.Contains(errMsg, "reset by peer") ||
							strings.Contains(errMsg, "request headers: small read buffer") ||
							strings.Contains(errMsg, "unexpected EOF") ||
							strings.Contains(errMsg, "i/o timeout") {
							brokenPipe = true
						}
					}
				}

				// 是否是 brokenPipe类型的错误，如果是，这里就不能往写入流中再写入内容
				// 如果是该类型的错误，就不需要返回任何数据给客户端
				// 代码参考gin recovery.go RecoveryWithWriter方法实现
				// If the connection is dead, we can't write a status to it.
				if brokenPipe {
					// ctx.Error(err.(error)) // nolint: errcheck
					ctx.AbortWithStatus(http.StatusInternalServerError)
					return
				}

				// 响应状态
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "server inner error",
				})
			}
		}()

		ctx.Next()
	}
}
