package handler

import (
	"encoding/json"
	"github.com/prynnekey/go-reggie/global"
	"net/http"
	"reggie_take_ut/internal/model"
	"reggie_take_ut/pkg/common"
	"strconv"
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
