package model

import (
	"math/big"
	"time"
)

// ShoppingCart 购物车实体
type ShoppingCart struct {
	ID         int64      `json:"id"`          // ID
	Name       string     `json:"name"`        // 名称
	UserID     int64      `json:"user_id"`     // 用户ID
	DishID     int64      `json:"dish_id"`     // 菜品ID
	SetmealID  int64      `json:"setmeal_id"`  // 套餐ID
	DishFlavor string     `json:"dish_flavor"` // 口味
	Number     int        `json:"number"`      // 数量
	Amount     *big.Float `json:"amount"`      // 金额
	Image      string     `json:"image"`       // 图片
	CreateTime time.Time  `json:"create_time"` // 创建时间
}
