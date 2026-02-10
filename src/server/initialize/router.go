package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/goer3/marinerx/common"
	"github.com/goer3/marinerx/middleware"
	"github.com/goer3/marinerx/router"
)

// 初始化路由
func Router() *gin.Engine {
	// 设置运行模式
	gin.SetMode(common.Config.System.Mode)

	// 创建一个没有中间件的 Gin 路由引擎
	r := gin.New()

	// 全局中间件
	r.Use(middleware.Exception)
	r.Use(middleware.AccessLogger)

	// 开放路由，无需认证
	OpenApiGroup := r.Group(common.SYSTEM_OPEN_API_PREFIX + "/" + common.SYSTEM_API_VERSION)
	router.OpenApiRouter(OpenApiGroup)

	// 开放路由，需要认证
	OpenApiWithAuthGroup := r.Group(common.SYSTEM_OPEN_API_PREFIX + "/" + common.SYSTEM_API_VERSION)
	router.OpenApiWithAuthRouter(OpenApiWithAuthGroup)

	return r
}
