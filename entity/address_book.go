package entity

import (
	"time"
)

// AddressBook 地址簿实体
type AddressBook struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`       // 用户id
	Consignee    string    `json:"consignee"`     // 收货人
	Phone        string    `json:"phone"`         // 手机号
	Sex          string    `json:"sex"`           // 性别 0 女 1 男
	ProvinceCode string    `json:"province_code"` // 省级区划编号
	ProvinceName string    `json:"province_name"` // 省级名称
	CityCode     string    `json:"city_code"`     // 市级区划编号
	CityName     string    `json:"city_name"`     // 市级名称
	DistrictCode string    `json:"district_code"` // 区级区划编号
	DistrictName string    `json:"district_name"` // 区级名称
	Detail       string    `json:"detail"`        // 详细地址
	Label        string    `json:"label"`         // 标签
	IsDefault    int       `json:"is_default"`    // 是否默认 0 否 1是
	CreateTime   time.Time `json:"create_time"`   // 创建时间
	UpdateTime   time.Time `json:"update_time"`   // 更新时间
	CreateUser   int64     `json:"create_user"`   // 创建人
	UpdateUser   int64     `json:"update_user"`   // 修改人
	IsDeleted    int       `json:"is_deleted"`    // 是否删除
}
