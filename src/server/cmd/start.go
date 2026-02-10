package cmd

import (
	"fmt"

	"github.com/goer3/marinerx/common"
	"github.com/goer3/marinerx/initialize"
	"github.com/goer3/marinerx/task"
	"github.com/spf13/cobra"
)

// 初始化命令
func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&common.ParamSystemListenHost, "host", "", common.ParamSystemListenHost, "指定服务启动监听地址")
	startCmd.Flags().IntVarP(&common.ParamSystemListenPort, "port", "", common.ParamSystemListenPort, "指定服务启动监听端口")
	startCmd.Flags().StringVarP(&common.ParamSystemConfigFile, "config", "", "", "指定服务启动配置文件")
	startCmd.Flags().IntVarP(&common.ParamSystemRoleLeader, "role-leader", "", common.ParamSystemRoleLeader, "指定服务是否参与领导者选举，0: 否，1: 是")
	startCmd.Flags().IntVarP(&common.ParamSystemRoleWorker, "role-worker", "", common.ParamSystemRoleWorker, "指定服务是否开启工作节点角色，0: 否，1: 是")
	startCmd.Flags().IntVarP(&common.ParamSystemRoleWeb, "role-web", "", common.ParamSystemRoleWeb, "指定服务是否开启 Web 后端服务角色，0: 否，1: 是")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 Mariner X 服务，更多参数请使用 --help 查看",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(common.LOGO)

		// 初始化配置
		initialize.Config()

		// 初始化日志
		initialize.SystemLogger()
		initialize.AccessLogger()

		// 判断是否是 Leader 节点
		if common.Config.System.Role.Leader {
			fmt.Println("启动为 Leader 节点...")
		}

		// 判断是否是 Worker 节点
		if common.Config.System.Role.Worker {
			fmt.Println("启动为 Worker 节点...")
		}

		// 判断是否是 Web 后端服务
		if common.Config.System.Role.Web {
			task.StartWebServer()
		} else {
			// 保持进程运行
			select {}
		}
	},
}
