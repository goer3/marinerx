package service

import (
	"errors"

	"github.com/dromara/carbon/v2"
	"github.com/goer3/marinerx/common"
	"github.com/goer3/marinerx/dto"
	"github.com/goer3/marinerx/model"
	"github.com/goer3/marinerx/pkg/utils"
	"github.com/jinzhu/copier"
)

// 创建用户
func SystemUserCreate(req *dto.SystemUserCreateRequest) error {
	// 验证过期时间格式
	var expireAt *carbon.Carbon
	if req.ExpireAt != "" {
		c := carbon.ParseByLayout(req.ExpireAt, common.TIME_SECOND, carbon.Local)
		if c == nil || c.IsInvalid() {
			return errors.New("过期时间格式不正确")
		}
		expireAt = c
	} else {
		expireAt = carbon.Now().AddYears(common.SYSTEM_USER_DEFAULT_EXPIRE_YEARS)
	}

	// 密码加密
	hashedPassword, err := utils.PasswordEncrypt(req.Password)
	if err != nil {
		return errors.New("密码加密失败：" + err.Error())
	}

	// 数据转换
	var user = model.SystemUser{
		Password: hashedPassword,
		ExpireAt: expireAt,
	}
	if err := copier.Copy(&user, &req); err != nil {
		common.SystemLog.Error("创建用户数据转换异常：", err.Error())
		return errors.New("创建用户数据转换异常：" + err.Error())
	}

	// 创建用户
	if err := common.DB.Create(&user).Error; err != nil {
		common.SystemLog.Error("创建用户失败：", err.Error())
		return errors.New("创建用户失败：" + err.Error())
	}

	return nil
}
