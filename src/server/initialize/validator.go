package initialize

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 手机号正则表达式
var mobileRegexp = regexp.MustCompile(`^1[3456789]\d{9}$`)

// 手机号验证
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	return mobileRegexp.MatchString(mobile)
}

// 初始化 gin 验证器
func Validator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证函数
		_ = v.RegisterValidation("mobile", ValidateMobile)
	}
}
