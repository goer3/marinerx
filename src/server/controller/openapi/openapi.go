package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goer3/marinerx/common"
	"github.com/goer3/marinerx/pkg/response"
)

// 健康检查处理器
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

// 版本信息处理器
func VersionHandler(ctx *gin.Context) {
	response.SuccessWithData(gin.H{
		"version": common.PROJECT_VERSION,
	})
}

// 系统信息处理器
func InformationHandler(ctx *gin.Context) {
	response.SuccessWithData(gin.H{
		"project": gin.H{
			"name":        common.PROJECT_NAME,
			"description": common.PROJECT_DESCRIPTION,
			"version":     common.PROJECT_VERSION,
		},
		"developer": gin.H{
			"name":       common.PROJECT_DEVELOPER,
			"email":      common.PROJECT_DEVELOPER_EMAIL,
			"repository": common.PROJECT_GITHUB_REPOSITORY,
		},
		"environment": gin.H{
			"go":    common.PROJECT_GO_VERSION,
			"react": common.PROJECT_REACT_VERSION,
			"antd":  common.PROJECT_ANTD_VERSION,
		},
	})
}
