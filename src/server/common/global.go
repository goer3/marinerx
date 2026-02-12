package common

import (
	"embed"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 项目常量
const (
	PROJECT_NAME              = "MARINER-X"
	PROJECT_DESCRIPTION       = "Mariner X，一站式云原生运维管理平台"
	PROJECT_VERSION           = "1.0.0"
	PROJECT_GO_VERSION        = "1.26.0+"
	PROJECT_REACT_VERSION     = "19+"
	PROJECT_ANTD_VERSION      = "6+"
	PROJECT_DEVELOPER         = "DK"
	PROJECT_DEVELOPER_EMAIL   = "1214966109@qq.com"
	PROJECT_GITHUB_REPOSITORY = "https://github.com/goer3/marinerx"
)

// 系统常量
const (
	SYSTEM_OPEN_API_PREFIX = "/openapi" // 系统开放API前缀
	SYSTEM_API_PREFIX      = "/api"     // 系统API前缀
	SYSTEM_API_VERSION     = "v1"       // 系统API版本
)

// 时间格式化
const (
	TIME_MILLISECOND string = "2006-01-02 15:04:05.000" // 毫秒时间格式化
	TIME_SECOND      string = "2006-01-02 15:04:05"     // 秒时间格式化
	TIME_DATE        string = "2006-01-02"              // 日期时间格式化
)

// 全局变量
var (
	FS        embed.FS           // 静态文件
	Config    *Configuration     // 配置文件解析保存数据
	SystemLog *zap.SugaredLogger // 系统日志
	AccessLog *zap.SugaredLogger // 访问日志
	DB        *gorm.DB           // MySQL 连接
	Cache     *redis.Client      // Redis 连接
)
