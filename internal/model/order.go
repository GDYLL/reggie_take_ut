package model

import (
	"math/big"
	"time"
)

// Orders 订单实体
type Orders struct {
	ID            int64      `json:"id"`              // 订单ID
	Number        string     `json:"number"`          // 订单号
	Status        int        `json:"status"`          // 订单状态 1待付款，2待派送，3已派送，4已完成，5已取消
	UserID        int64      `json:"user_id"`         // 下单用户ID
	AddressBookID int64      `json:"address_book_id"` // 地址ID
	OrderTime     time.Time  `json:"order_time"`      // 下单时间
	CheckoutTime  time.Time  `json:"checkout_time"`   // 结账时间
	PayMethod     int        `json:"pay_method"`      // 支付方式 1微信，2支付宝
	Amount        *big.Float `json:"amount"`          // 实收金额
	Remark        string     `json:"remark"`          // 备注
	UserName      string     `json:"user_name"`       // 用户名
	Phone         string     `json:"phone"`           // 手机号
	Address       string     `json:"address"`         // 地址
	Consignee     string     `json:"consignee"`       // 收货人
}

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

// OrdersDto 扩展的订单实体
type OrdersDto struct {
	Orders
	UserName     string        `json:"user_name"`     // 用户名
	Phone        string        `json:"phone"`         // 手机号
	Address      string        `json:"address"`       // 地址
	Consignee    string        `json:"consignee"`     // 收货人
	OrderDetails []OrderDetail `json:"order_details"` // 订单明细列表
}
