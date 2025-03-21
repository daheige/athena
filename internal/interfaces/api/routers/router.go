package routers

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/daheige/athena/internal/application"
	"github.com/daheige/athena/internal/interfaces/api/handlers"
	"github.com/daheige/athena/internal/interfaces/api/middleware"
	"github.com/daheige/athena/internal/providers"
)

// InitRouters 初始化router规则
func InitRouters(router *gin.Engine, repos *providers.Repositories) {
	// 访问日志中间件处理
	logWare := &middleware.LogWare{}

	// 对所有的请求进行性能监控，一般来说生产环境，可以对指定的接口做性能监控
	router.Use(logWare.Access(), logWare.Recover(), middleware.TimeoutHandler(10*time.Second))

	// gin 框架prometheus接入
	router.Use(middleware.WrapMonitor())

	// 路由找不到的情况
	router.NoRoute(middleware.NotFoundHandler())

	// 创建user service
	userService := application.NewUserService(repos.UserRepo, repos.UserCache)
	indexHandler := handlers.NewIndexHandler(userService)
	router.GET("/", indexHandler.Home)

	apiGroup := router.Group("api") // 定义路由组
	apiGroup.GET("foo", indexHandler.Foo)
	apiGroup.GET("user", indexHandler.User)
	apiGroup.POST("users", indexHandler.BatchUsers)

	// 其他路由请自行添加，如果路由规则比较多，请自行拆分到不同的文件中即可
}
