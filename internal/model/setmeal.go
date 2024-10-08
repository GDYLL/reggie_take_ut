package model

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

// SetmealDish 套餐菜品关系实体
type SetmealDish struct {
	ID         int64      `json:"id"`          // ID
	SetmealID  int64      `json:"setmeal_id"`  // 套餐ID
	DishID     int64      `json:"dish_id"`     // 菜品ID
	Name       string     `json:"name"`        // 菜品名称 （冗余字段）
	Price      *big.Float `json:"price"`       // 菜品原价
	Copies     int        `json:"copies"`      // 份数
	Sort       int        `json:"sort"`        // 排序
	CreateTime time.Time  `json:"create_time"` // 创建时间
	UpdateTime time.Time  `json:"update_time"` // 更新时间
	CreateUser int64      `json:"create_user"` // 创建人
	UpdateUser int64      `json:"update_user"` // 修改人
	IsDeleted  int        `json:"is_deleted"`  // 是否删除
}

// SetmealDto 扩展的套餐实体
type SetmealDto struct {
	Setmeal
	SetmealDishes []SetmealDish `json:"setmeal_dishes"` // 套餐菜品关系列表
	CategoryName  string        `json:"category_name"`  // 分类名称
}
