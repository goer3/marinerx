package common

// 配置解析入口
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" yaml:"system" json:"system"`
	Log    LogConfiguration    `mapstructure:"log" yaml:"log" json:"log"`
}

// 系统配置
type SystemConfiguration struct {
	Mode   string              `mapstructure:"mode" yaml:"mode" json:"mode"`
	Listen ListenConfiguration `mapstructure:"listen" yaml:"listen" json:"listen"`
	Role   RoleConfiguration   `mapstructure:"role" yaml:"role" json:"role"`
}

// 监听配置
type ListenConfiguration struct {
	Host string `mapstructure:"host" yaml:"host" json:"host"`
	Port int    `mapstructure:"port" yaml:"port" json:"port"`
}

// 角色配置
type RoleConfiguration struct {
	Leader bool `mapstructure:"leader" yaml:"leader" json:"leader"`
	Worker bool `mapstructure:"worker" yaml:"worker" json:"worker"`
	Web    bool `mapstructure:"web" yaml:"web" json:"web"`
}

// 日志配置
type LogConfiguration struct {
	System BaseLogConfiguration `mapstructure:"system" yaml:"system" json:"system"`
	Access BaseLogConfiguration `mapstructure:"access" yaml:"access" json:"access"`
}

// 基础日志配置
type BaseLogConfiguration struct {
	Enabled   bool                    `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	Level     string                  `mapstructure:"level" yaml:"level" json:"level"`
	Formatter string                  `mapstructure:"formatter" yaml:"formatter" json:"formatter"`
	Path      string                  `mapstructure:"path" yaml:"path" json:"path"`
	Prefix    string                  `mapstructure:"prefix" yaml:"prefix" json:"prefix"`
	Rolling   LogRollingConfiguration `mapstructure:"rolling" yaml:"rolling" json:"rolling"`
}

// 日志切割配置
type LogRollingConfiguration struct {
	Enabled    bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	MaxSize    int  `mapstructure:"max_size" yaml:"max_size" json:"max_size"`
	MaxBackups int  `mapstructure:"max_backups" yaml:"max_backups" json:"max_backups"`
	MaxAge     int  `mapstructure:"max_age" yaml:"max_age" json:"max_age"`
	Compress   bool `mapstructure:"compress" yaml:"compress" json:"compress"`
}
