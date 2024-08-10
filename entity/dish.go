package entity

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
