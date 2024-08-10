package dto

import (
	"reggie_take_ut/entity"
)

// DishDto 扩展的菜品实体
type DishDto struct {
	entity.Dish
	Flavors      []entity.DishFlavor `json:"flavors"`      // 菜品口味列表
	CategoryName string              `json:"categoryName"` // 分类名称
	Copies       int                 `json:"copies"`       // 份数
}
