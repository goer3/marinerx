package initialize

import (
	"github.com/goer3/marinerx/model"
)

// 菜单初始化
var SystemMenus = []model.SystemMenu{}

// 角色初始化
var SystemRoles = []model.SystemRole{}

// 用户初始化
var SystemUsers = []model.SystemUser{}

// 接口初始化
var SystemApiCategories = []model.SystemApiCategory{
	{
		Id:   10000,
		Name: "对外开放接口",
		SystemApis: []model.SystemApi{
			{Id: 10001, Name: "健康检查", Method: "GET", Api: "/openapi/v1/health", IsOpenApi: 1, SystemApiCategoryId: 10000},
			{Id: 10002, Name: "版本信息", Method: "GET", Api: "/openapi/v1/version", IsOpenApi: 1, SystemApiCategoryId: 10000},
			{Id: 10003, Name: "系统信息", Method: "GET", Api: "/openapi/v1/information", IsOpenApi: 1, SystemApiCategoryId: 10000},
		},
	},
	{
		Id:   11000,
		Name: "免授权接口",
		SystemApis: []model.SystemApi{
			{Id: 11001, Name: "登录（用户）", Method: "GET", Api: "/api/v1/login", IsOpenApi: 1, SystemApiCategoryId: 11000},
			{Id: 11002, Name: "登录（钉钉）", Method: "GET", Api: "/api/v1/login/dingtalk", IsOpenApi: 1, SystemApiCategoryId: 11000},
			{Id: 11003, Name: "登录（飞书）", Method: "GET", Api: "/api/v1/login/feishu", IsOpenApi: 1, SystemApiCategoryId: 11000},
			{Id: 11004, Name: "登录（企业微信）", Method: "GET", Api: "/api/v1/login/wechat", IsOpenApi: 1, SystemApiCategoryId: 11000},
			{Id: 11005, Name: "登出", Method: "GET", Api: "/api/v1/logout", IsOpenApi: 1, SystemApiCategoryId: 11000},
		},
	},
	{
		Id:   12000,
		Name: "用户模块",
		SystemApis: []model.SystemApi{
			{Id: 12001, Name: "创建用户", Method: "POST", Api: "/api/v1/system/user/create", IsOpenApi: 0, SystemApiCategoryId: 12000},
			{Id: 12002, Name: "创建用户（批量）", Method: "POST", Api: "/api/v1/system/user/create/batch", IsOpenApi: 0, SystemApiCategoryId: 12000},
			{Id: 12003, Name: "获取用户列表", Method: "GET", Api: "/api/v1/system/user/list", IsOpenApi: 0, SystemApiCategoryId: 12000},
			{Id: 12004, Name: "获取用户详情", Method: "GET", Api: "/api/v1/system/user/detail", IsOpenApi: 0, SystemApiCategoryId: 12000},
		},
	},
	{
		Id:         13000,
		Name:       "角色模块",
		SystemApis: []model.SystemApi{},
	},
	{
		Id:         14000,
		Name:       "菜单模块",
		SystemApis: []model.SystemApi{},
	},
	{
		Id:         15000,
		Name:       "接口类型模块",
		SystemApis: []model.SystemApi{},
	},
	{
		Id:         16000,
		Name:       "接口模块",
		SystemApis: []model.SystemApi{},
	},
}

// 同步接口数据，需要如果不存在，则添加
// func SyncSystemApiData() {
// 	for _, category := range SystemApiCategories {
// 		// 判断类似是否存在，如果不存在，则添加，如果存在，则判断内容是否一致，如果不一致，则更新

// 		for _, api := range category.SystemApis {
// 			// 判断接口是否存在，如果不存在，则添加，如果存在，则判断内容是否一致，如果不一致，则更新
// 		}
// 	}
// }
