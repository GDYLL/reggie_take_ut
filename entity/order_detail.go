package entity

import (
	"math/big"
)

// OrderDetail 订单明细实体
type OrderDetail struct {
	ID         int64      `json:"id"`          // 订单明细ID
	Name       string     `json:"name"`        // 名称
	OrderID    int64      `json:"order_id"`    // 订单ID
	DishID     int64      `json:"dish_id"`     // 菜品ID
	SetmealID  int64      `json:"setmeal_id"`  // 套餐ID
	DishFlavor string     `json:"dish_flavor"` // 口味
	Number     int        `json:"number"`      // 数量
	Amount     *big.Float `json:"amount"`      // 金额
	Image      string     `json:"image"`       // 图片
}
