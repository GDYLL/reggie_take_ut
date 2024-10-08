package model

// User 用户信息实体
type User struct {
	ID       int64  `json:"id"`        // ID
	Name     string `json:"name"`      // 姓名
	Phone    string `json:"phone"`     // 手机号
	Sex      string `json:"sex"`       // 性别 0 女 1 男
	IDNumber string `json:"id_number"` // 身份证号
	Avatar   string `json:"avatar"`    // 头像
	Status   int    `json:"status"`    // 状态 0:禁用，1:正常
}
