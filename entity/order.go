package entity

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
