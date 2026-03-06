package initialize

import (
	"os"

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
	r.Use(middleware.AccessLogger)
	r.Use(middleware.Exception)

	// JWT 中间件
	auth, err := middleware.JWTAuth()
	if err != nil {
		common.SystemLog.Fatal("JWT 中间件初始化异常：" + err.Error())
		os.Exit(1)
	}

	// 开放 API 前缀
	OpenApiPrefix := common.SYSTEM_OPEN_API_PREFIX + "/" + common.SYSTEM_API_VERSION

	// API 前缀
	ApiPrefix := common.SYSTEM_API_PREFIX + "/" + common.SYSTEM_API_VERSION

	{
		// 开放路由，无需认证
		OpenApiGroup := r.Group(OpenApiPrefix)
		router.OpenApiRouter(OpenApiGroup)
	}
	{
		// 开放路由，需要认证
		OpenApiWithAuthGroup := r.Group(ApiPrefix)
		router.OpenApiWithAuthRouter(OpenApiWithAuthGroup)
	}
	{
		// 基础路由，无需认证
		ApiGroup := r.Group(ApiPrefix)
		router.ApiRouter(ApiGroup, auth)
	}
	{
		// 系统模块路由，需要认证
		SystemApiGroup := ApiPrefix + "/system"
		// 用户模块路由，需要认证
		router.SystemUserApiWithAuthRouter(r.Group(SystemApiGroup+"/user"), auth)
	}

	return r
}
