package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goer3/marinerx/dto"
	"github.com/goer3/marinerx/pkg/response"
	"github.com/goer3/marinerx/pkg/utils"
	"github.com/goer3/marinerx/service"
)

// 创建用户
func SystemUserCreateHandler(ctx *gin.Context) {
	var req dto.SystemUserCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := utils.GetValidateErrorMessage(err, &req)
		response.FailureWithMessage(msg)
		return
	}

	if err := service.SystemUserCreate(&req); err != nil {
		response.FailureWithMessage(err.Error())
		return
	}

	response.SuccessWithMessage("用户创建成功")
}

// 批量创建用户
func SystemUserBatchCreateHandler(ctx *gin.Context) {}

// 获取用户列表
func SystemUserListHandler(ctx *gin.Context) {}

// 获取用户详情
func SystemUserDetailHandler(ctx *gin.Context) {}
