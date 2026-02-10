package common

// 参数初始化配置，用于命令行参数初始化
var (
	// 服务默认监听地址
	ParamSystemListenHost = ""
	// 服务默认监听端口
	ParamSystemListenPort = -1
	// 服务默认配置文件，默认空字符串，不能设置默认值
	ParamSystemConfigFile = ""
	// 服务是否参与领导者选举, 0: 不参与, 1: 参与
	ParamSystemRoleLeader = -1
	// 服务是否开启工作节点角色, 0: 不开启, 1: 开启
	ParamSystemRoleWorker = -1
	// 服务是否开启Web后端服务角色, 0: 不开启, 1: 开启
	ParamSystemRoleWeb = -1
)
