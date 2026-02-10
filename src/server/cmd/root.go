package cmd

import (
	"github.com/goer3/marinerx/common"
	"github.com/spf13/cobra"
)

// 命令入口
var rootCmd = &cobra.Command{
	Use:   "marinerx",
	Short: common.PROJECT_DESCRIPTION,
}

func Execute() {
	rootCmd.Execute()
}
