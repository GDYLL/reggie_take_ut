package handler

import (
	"encoding/json"
	"github.com/prynnekey/go-reggie/global"
	"net/http"
	"reggie_take_ut/internal/model"
	"reggie_take_ut/pkg/common"
	"reggie_take_ut/pkg/session"
	"strconv"
	"time"
)

type CategoryController struct {
}

func (c CategoryController) Page(w http.ResponseWriter, r *http.Request) {

	num := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	pageNum, _ := strconv.Atoi(num)
	if pageNum <= 0 {
		pageNum = 1
	}

	pageSize, _ := strconv.Atoi(size)
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (pageNum - 1) * pageSize

	var category []model.Category
	if global.DB.Table("category").Offset(offset).Limit(pageSize).Find(&category).Error != nil {
		http.Error(w, "查询失败", http.StatusInternalServerError)
		return
	}
	var total int64
	if global.DB.Table("category").Count(&total).Error != nil {
		http.Error(w, "查询失败", http.StatusInternalServerError)
		return
	}
	responseData := model.ResponseData{
		Records: category,
		Total:   total,
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(common.Result{}.Success(responseData))
	if err != nil {
		http.Error(w, "JSON 编码失败", http.StatusInternalServerError)
		return
	}
	return
}

func (c CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	var category model.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "JSON 解码失败", http.StatusBadRequest)
		return
	}

	// 从 session 中获取当前用户 ID
	session, err1 := session.Store.Get(r, session.SessionName)
	if err1 != nil {
		http.Error(w, "获取用户信息失败", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// 查询当前套餐或菜品是否已存在
	var count int64
	if global.DB.Model(&model.Category{}).Where("name = ?", category.Name).Count(&count); count > 0 {
		json.NewEncoder(w).Encode(common.Result{}.Error("该分类已存在"))
		return
	}

	category.UpdateUser = session.Values["employee_id"].(string)
	category.UpdateTime = time.Now()

	if err := global.DB.Model(&model.Category{}).Omit("id").Where("id = ?", category.ID).Updates(&category).Error; err != nil {
		http.Error(w, "更新失败", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(common.Result{}.Success("更新成功"))
	if err != nil {
		http.Error(w, "系统错误", http.StatusInternalServerError)
		return
	}
	return
}

func (c CategoryController) Save(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "JSON 解码失败", http.StatusBadRequest)
		return
	}
	category.CreateTime = time.Now()
	category.UpdateTime = time.Now()

	// 从 session 中获取当前用户 ID
	session, err1 := session.Store.Get(r, session.SessionName)
	if err1 != nil {
		http.Error(w, "获取用户信息失败", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// 查询当前套餐或菜品是否已存在
	var count int64
	if global.DB.Model(&model.Category{}).Where("name = ?", category.Name).Count(&count); count > 0 {
		json.NewEncoder(w).Encode(common.Result{}.Error("该分类已存在"))
		return
	}

	employeeID, _ := session.Values["employee_id"].(string)
	category.CreateUser = employeeID
	category.UpdateUser = employeeID

	if err := global.DB.Create(&category).Error; err != nil {
		http.Error(w, "创建失败", http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(common.Result{}.Success("创建成功"))
	if err != nil {
		http.Error(w, "系统错误", http.StatusInternalServerError)
	}
	return
}
