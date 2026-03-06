package router

import (
	jwt "github.com/appleboy/gin-jwt/v3"
	"github.com/gin-gonic/gin"
	"github.com/goer3/marinerx/controller/api"
)

// 基础路由，无需登录
func ApiRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.POST("/login", auth.LoginHandler)          // 登录（用户）
	rg.POST("/login/dingtalk", auth.LoginHandler) // 登录（钉钉）
	rg.POST("/login/feishu", auth.LoginHandler)   // 登录（飞书）
	rg.POST("/login/wechat", auth.LoginHandler)   // 登录（企业微信）
	return rg
}

// 用户路由，需要登录
func SystemUserApiWithAuthRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.POST("/create", api.SystemUserCreateHandler)            // 用户创建
	rg.POST("/create/batch", api.SystemUserBatchCreateHandler) // 用户创建（批量）
	rg.GET("/list", api.SystemUserListHandler)                 // 用户列表
	rg.GET("/detail", api.SystemUserDetailHandler)             // 用户详情
	return rg
}
