package entity

import "time"

// Category 分类实体
type Category struct {
	ID         int64     `json:"id"`         // 分类ID
	Type       string    `json:"type"`       // 类型 1 菜品分类 2 套餐分类
	Name       string    `json:"name"`       // 分类名称
	Sort       string    `json:"sort"`       // 顺序
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
	CreateUser int64     `json:"createUser"` // 创建人
	UpdateUser int64     `json:"updateUser"` // 修改人
	//IsDeleted  int       `json:"is_deleted"`  // 是否删除
}
