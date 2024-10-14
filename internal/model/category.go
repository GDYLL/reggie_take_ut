package model

import "time"

// Category 分类实体
type Category struct {
	ID         string    `gorm:"<-:update" json:"id"`           // 分类ID
	Type       string    `json:"type"`                          // 类型 1 菜品分类 2 套餐分类
	Name       string    `json:"name"`                          // 分类名称
	Sort       string    `json:"sort"`                          // 顺序
	CreateTime time.Time `gorm:"crete_time" json:"createTime"`  // 创建时间
	UpdateTime time.Time `gorm:"update_time" json:"updateTime"` // 更新时间
	CreateUser string    `gorm:"create_user" json:"createUser"` // 创建人
	UpdateUser string    `gorm:"update_user" json:"updateUser"` // 修改人
}
