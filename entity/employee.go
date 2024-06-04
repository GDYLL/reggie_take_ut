package entity

import (
	"time"
)

// Employee 员工实体
type Employee struct {
	ID         int64     `json:"id"`          // 员工ID
	Username   string    `json:"username"`    // 用户名
	Name       string    `json:"name"`        // 姓名
	Password   string    `json:"password"`    // 密码
	Phone      string    `json:"phone"`       // 手机号
	Sex        string    `json:"sex"`         // 性别
	IDNumber   string    `json:"id_number"`   // 身份证号码
	Status     int       `json:"status"`      // 状态
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	CreateUser int64     `json:"create_user"` // 创建人
	UpdateUser int64     `json:"update_user"` // 修改人
}
