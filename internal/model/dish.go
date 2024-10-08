package model

import (
	"time"
)

// Dish 菜品实体
type Dish struct {
	ID          int64     `json:"id"`          // 菜品ID
	Name        string    `json:"name"`        // 菜品名称
	CategoryId  int64     `json:"categoryId"`  // 菜品分类ID
	Price       float32   `json:"price"`       // 菜品价格
	Code        string    `json:"code"`        // 商品码
	Image       string    `json:"image"`       // 图片
	Description string    `json:"description"` // 描述信息
	Status      int       `json:"status"`      // 0 停售 1 起售
	Sort        int       `json:"sort"`        // 顺序
	CreateTime  time.Time `json:"createTime"`  // 创建时间
	UpdateTime  time.Time `json:"updateTime"`  // 更新时间
	CreateUser  int64     `json:"createUser"`  // 创建人
	UpdateUser  int64     `json:"updateUser"`  // 修改人
}

// DishFlavor 菜品口味实体
type DishFlavor struct {
	ID         int64     `json:"id"`          // 菜品口味ID
	DishID     int64     `json:"dish_id"`     // 菜品ID
	Name       string    `json:"name"`        // 口味名称
	Value      string    `json:"value"`       // 口味数据列表
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	CreateUser int64     `json:"create_user"` // 创建人
	UpdateUser int64     `json:"update_user"` // 修改人
	IsDeleted  int       `json:"is_deleted"`  // 是否删除
}

// DishDto 扩展的菜品实体
type DishDto struct {
	Dish
	Flavors      []DishFlavor `json:"flavors"`      // 菜品口味列表
	CategoryName string       `json:"categoryName"` // 分类名称
	Copies       int          `json:"copies"`       // 份数
}
