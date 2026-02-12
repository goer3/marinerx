package initialize

import "github.com/goer3/marinerx/common"

// 数据表结构同步
func Table() {
	models := []any{
		// &model.SystemRole{},
		// &model.SystemUser{},
	}

	for _, m := range models {
		tableName := ""
		// 获取表名
		if t, ok := m.(interface{ TableName() string }); ok {
			tableName = t.TableName()
		} else {
			tableName = "unknown"
		}
		common.SystemLog.Info("开始同步数据表：" + tableName)
		if err := common.DB.AutoMigrate(m); err != nil {
			common.SystemLog.Error("同步数据表失败：" + tableName + "，错误：" + err.Error())
		}
	}
	common.SystemLog.Info("数据表同步完成")
}
