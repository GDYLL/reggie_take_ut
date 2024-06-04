package entity

import (
	"math/big"
	"time"
)

// Dish 菜品实体
type Dish struct {
	ID          int64      `json:"id"`          // 菜品ID
	Name        string     `json:"name"`        // 菜品名称
	CategoryID  int64      `json:"category_id"` // 菜品分类ID
	Price       *big.Float `json:"price"`       // 菜品价格
	Code        string     `json:"code"`        // 商品码
	Image       string     `json:"image"`       // 图片
	Description string     `json:"description"` // 描述信息
	Status      int        `json:"status"`      // 0 停售 1 起售
	Sort        int        `json:"sort"`        // 顺序
	CreateTime  time.Time  `json:"create_time"` // 创建时间
	UpdateTime  time.Time  `json:"update_time"` // 更新时间
	CreateUser  int64      `json:"create_user"` // 创建人
	UpdateUser  int64      `json:"update_user"` // 修改人
	IsDeleted   int        `json:"is_deleted"`  // 是否删除
}
