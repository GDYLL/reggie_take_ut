package entity

import (
	"math/big"
	"time"
)

// Setmeal 套餐实体
type Setmeal struct {
	ID          int64      `json:"id"`          // 套餐ID
	CategoryID  int64      `json:"category_id"` // 分类ID
	Name        string     `json:"name"`        // 套餐名称
	Price       *big.Float `json:"price"`       // 套餐价格
	Status      int        `json:"status"`      // 状态 0:停用 1:启用
	Code        string     `json:"code"`        // 编码
	Description string     `json:"description"` // 描述信息
	Image       string     `json:"image"`       // 图片
	CreateTime  time.Time  `json:"create_time"` // 创建时间
	UpdateTime  time.Time  `json:"update_time"` // 更新时间
	CreateUser  int64      `json:"create_user"` // 创建人
	UpdateUser  int64      `json:"update_user"` // 修改人
	IsDeleted   int        `json:"is_deleted"`  // 是否删除
}
