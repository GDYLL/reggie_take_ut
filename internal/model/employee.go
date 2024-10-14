package model

import (
	"time"
)

// Employee 员工实体
type Employee struct {
	ID         string    `gorm:"<-:update" json:"id"`           // 员工ID
	Username   string    `json:"username"`                      // 用户名
	Name       string    `json:"name"`                          // 姓名
	Password   string    `json:"password"`                      // 密码
	Phone      string    `json:"phone"`                         // 手机号
	Sex        string    `json:"sex"`                           // 性别
	IDNumber   string    `gorm:"id_number" json:"idNumber"`     // 身份证号码
	Status     *int      `json:"status"`                        // 状态,指针避免 gorm 的默认值不更新
	CreateTime time.Time `gorm:"cerate_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"update_time" json:"updateTime"` // 更新时间
	CreateUser int64     `gorm:"create_user" json:"createUser"` // 创建人
	UpdateUser int64     `gorm:"update_user" json:"updateUser"` // 修改人
}
