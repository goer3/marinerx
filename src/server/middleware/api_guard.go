package middleware

import (
	"github.com/gin-gonic/gin"
)

// API 接口授权中间件
func ApiGuard(ctx *gin.Context) {
	// 用户角色关键字
	// 请求路径，对于带参数的 GET 请求需要去掉参数
	// 请求方法
	// 判断角色是否在免授权角色中
	ctx.Next()
}
