package cmd

import (
	"github.com/goer3/marinerx/common"
	"github.com/goer3/marinerx/initialize"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateTableCmd)
	migrateTableCmd.Flags().StringVarP(&common.ParamSystemConfigFile, "config", "", "", "指定服务启动配置文件")
}

// 数据迁移
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据同步，支持数据表同步和基础数据初始化",
}

// 数据表迁移
var migrateTableCmd = &cobra.Command{
	Use:   "table",
	Short: "数据表结构同步",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Config()
		initialize.SystemLogger()
		initialize.MySQL()
		initialize.Table()
	},
}
