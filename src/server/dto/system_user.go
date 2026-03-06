package dto

// binding 内置的验证包括：
// 不能为空，并且不能没有这个字段
// required：必填字段，如：binding:"required"

// 针对字符串的长度
// min 最小长度，如：binding:"min=5"
// max 最大长度，如：binding:"max=10"
// len 长度，如：binding:"len=6"

// 针对数字的大小
// eq 等于，如：binding:"eq=3"
// ne 不等于，如：binding:"ne=12"
// gt 大于，如：binding:"gt=10"
// gte 大于等于，如：binding:"gte=10"
// lt 小于，如：binding:"lt=10"
// lte 小于等于，如：binding:"lte=10"

// 针对同级字段的
// eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`
// nefield 不等于其他字段的值，如：PassWord string `binding:"nefield=Password"`

// - 忽略字段，如：binding:"-"

// 枚举，如：oneof=red green

// 字符串
// contains=fengfeng  // 包含fengfeng的字符串
// excludes // 不包含
// startswith  // 字符串前缀
// endswith  // 字符串后缀

// 数组
// dive 后面的验证就是针对数组中的每一个元素

// 网络验证
// ip
// ipv4
// ipv6
// uri
// url

// 日期验证
// datetime=2006-01-02

// 用户创建请求 - 必填字段用非指针类型
type SystemUserCreateRequest struct {
	Nickname     string `json:"nickname" binding:"required,min=2,max=50" msg:"昵称不能为空，长度必须在2-50之间"`
	Username     string `json:"username" binding:"required,min=2,max=50" msg:"用户名不能为空，长度必须在2-50之间"`
	Password     string `json:"password" binding:"required,min=6,max=128" msg:"密码不能为空，长度必须在6-128之间"`
	Mobile       string `json:"mobile" binding:"required,mobile" msg:"手机号格式不正确"`
	HideMobile   uint   `json:"hide_mobile" binding:"required,oneof=0 1" msg:"隐藏手机号字段不能为空，且只能为0或1"`
	Email        string `json:"email" binding:"required,email" msg:"邮箱格式不正确"`
	Gender       uint   `json:"gender" binding:"required,oneof=0 1 2" msg:"性别字段不能为空，且只能为0（未知）、1（男）或2（女）"`
	AvatarUrl    string `json:"avatar_url" binding:"url" msg:"头像URL格式不正确"`
	Status       uint   `json:"status" binding:"required,oneof=0 1" msg:"状态字段不能为空，且只能为0（禁用）或1（启用）"`
	ExpireAt     string `json:"expire_at" binding:"datetime=2006-01-02 15:04:05" msg:"过期时间格式不正确，必须为2006-01-02 15:04:05格式"`
	SystemRoleId uint   `json:"system_role_id" binding:"required" msg:"系统角色ID不能为空"`
}
