package initialize

import (
	"bytes"
	"log"
	"net"
	"os"

	"github.com/goer3/marinerx/common"
	"github.com/spf13/viper"
)

// 配置初始化
func Config() {
	var bs []byte
	var err error
	var configType string = "yaml"
	var configFile string = "config/development.yaml"

	// 初始化 viper，设置配置文件类型
	v := viper.New()
	v.SetConfigType(configType)

	// 读取配置文件内容
	if common.ParamSystemConfigFile != "" {
		// 如果启动参数有指定配置文件，则使用指定的外部配置文件进行初始化
		log.Printf("使用外部配置文件进行初始化: %s", common.ParamSystemConfigFile)
		configFile = common.ParamSystemConfigFile
		bs, err = os.ReadFile(configFile)
	} else {
		// 否则使用默认配置文件，读取方式不一样
		log.Printf("使用默认配置文件进行初始化: %s", configFile)
		bs, err = common.FS.ReadFile(configFile)
	}

	// 读取配置文件失败
	if err != nil {
		log.Fatalln("读取配置文件失败：" + err.Error())
	}

	// viper 解析配置文件
	if err = v.ReadConfig(bytes.NewReader(bs)); err != nil {
		log.Fatalln("解析配置文件失败：" + err.Error())
	}

	// 反序列化配置，使用 UnmarshalExact 进行严格类型检查
	if err = v.UnmarshalExact(&common.Config); err != nil {
		log.Fatalln("反序列化配置失败（类型不匹配）：" + err.Error())
	}

	// 命令行参数解析

	// 监听地址校验: --host
	if common.ParamSystemListenHost != "" {
		common.Config.System.Listen.Host = common.ParamSystemListenHost
	}
	if net.ParseIP(common.Config.System.Listen.Host) == nil {
		log.Fatalln("监听地址参数 --host 设置错误，仅支持 IPV4 地址")
	}

	// 监听端口校验: --port
	if common.ParamSystemListenPort != -1 {
		common.Config.System.Listen.Port = common.ParamSystemListenPort
	}
	if common.Config.System.Listen.Port < 80 || common.Config.System.Listen.Port > 65535 {
		log.Fatalln("监听端口参数 --port 设置错误，仅支持 80-65535 之间的整数")
	}

	// Leader 角色: --role-leader
	if common.ParamSystemRoleLeader != -1 {
		switch common.ParamSystemRoleLeader {
		case 0:
			common.Config.System.Role.Leader = false
		case 1:
			common.Config.System.Role.Leader = true
		default:
			log.Fatalln("角色参数 --role-leader 设置错误，仅支持 0 或者 1")
		}
	}

	// Worker 角色: --role-worker
	if common.ParamSystemRoleWorker != -1 {
		switch common.ParamSystemRoleWorker {
		case 0:
			common.Config.System.Role.Worker = false
		case 1:
			common.Config.System.Role.Worker = true
		default:
			log.Fatalln("角色参数 --role-worker 设置错误，仅支持 0 或者 1")
		}
	}

	// Web 角色: --role-web
	if common.ParamSystemRoleWeb != -1 {
		switch common.ParamSystemRoleWeb {
		case 0:
			common.Config.System.Role.Web = false
		case 1:
			common.Config.System.Role.Web = true
		default:
			log.Fatalln("角色参数 --role-web 设置错误，仅支持 0 或者 1")
		}
	}

	// 如果所有角色都是 False，则退出
	if !common.Config.System.Role.Leader && !common.Config.System.Role.Worker && !common.Config.System.Role.Web {
		log.Fatalln("所有角色都未启用，无法启动服务，服务即将退出...")
	}
}
