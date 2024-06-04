package dto

import (
	"reggie_take_ut/entity"
)

// SetmealDto 扩展的套餐实体
type SetmealDto struct {
	entity.Setmeal
	SetmealDishes []entity.SetmealDish `json:"setmeal_dishes"` // 套餐菜品关系列表
	CategoryName  string               `json:"category_name"`  // 分类名称
}
