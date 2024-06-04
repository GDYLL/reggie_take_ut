package entity

import (
	"time"
)

// DishFlavor 菜品口味实体
type DishFlavor struct {
	ID         int64     `json:"id"`          // 菜品口味ID
	DishID     int64     `json:"dish_id"`     // 菜品ID
	Name       string    `json:"name"`        // 口味名称
	Value      string    `json:"value"`       // 口味数据列表
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	CreateUser int64     `json:"create_user"` // 创建人
	UpdateUser int64     `json:"update_user"` // 修改人
	IsDeleted  int       `json:"is_deleted"`  // 是否删除
}
