package entity

import (
	"time"
)

// Category 分类实体
type Category struct {
	ID         int64     `json:"id"`          // 分类ID
	Type       int       `json:"type"`        // 类型 1 菜品分类 2 套餐分类
	Name       string    `json:"name"`        // 分类名称
	Sort       int       `json:"sort"`        // 顺序
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	CreateUser int64     `json:"create_user"` // 创建人
	UpdateUser int64     `json:"update_user"` // 修改人
	IsDeleted  int       `json:"is_deleted"`  // 是否删除
}
