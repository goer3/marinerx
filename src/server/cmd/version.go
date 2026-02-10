package cmd

import (
	"fmt"

	"github.com/goer3/marinerx/common"
	"github.com/spf13/cobra"
)

// 初始化命令
func init() {
	rootCmd.AddCommand(versionCmd)
}

// 版本命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示 Mariner X 当前版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(common.PROJECT_VERSION)
	},
}
