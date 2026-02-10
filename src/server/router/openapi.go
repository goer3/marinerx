package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goer3/marinerx/controller/openapi"
)

// 开放路由，无需认证
func OpenApiRouter(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/health", openapi.HealthHandler)           // 健康检查
	rg.GET("/version", openapi.VersionHandler)         // 版本信息
	rg.GET("/information", openapi.InformationHandler) // 系统信息
	return rg
}

// 开放路由，需要认证
func OpenApiWithAuthRouter(rg *gin.RouterGroup) gin.IRoutes {
	return rg
}
