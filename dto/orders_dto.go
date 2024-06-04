package dto

import (
	"reggie_take_ut/entity"
)

// OrdersDto 扩展的订单实体
type OrdersDto struct {
	entity.Orders
	UserName     string               `json:"user_name"`     // 用户名
	Phone        string               `json:"phone"`         // 手机号
	Address      string               `json:"address"`       // 地址
	Consignee    string               `json:"consignee"`     // 收货人
	OrderDetails []entity.OrderDetail `json:"order_details"` // 订单明细列表
}
