package handler

import (
	"encoding/json"
	"github.com/prynnekey/go-reggie/global"
	"gorm.io/gorm"
	"net/http"
	"reggie_take_ut/internal/model"
	"reggie_take_ut/pkg/common"
	"strconv"
)

type DishController struct {
}

func (dc DishController) Page(w http.ResponseWriter, r *http.Request) {
	pageNum, pageSize := getPaginationParams(r)
	offset := (pageNum - 1) * pageSize

	var dishes []model.Dish
	var total int64
	var err error

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("dish").Count(&total).Error; err != nil {
			return err
		}
		if err := tx.Table("dish").Offset(offset).Limit(pageSize).Find(&dishes).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		http.Error(w, "查询失败", http.StatusInternalServerError)
	}

	dishDtos, err := getDishDtos(dishes)
	if err != nil {
		http.Error(w, "查询失败", http.StatusInternalServerError)
		return
	}

	responseData := model.ResponseData{
		Records: dishDtos,
		Total:   total,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(common.Result{}.Success(responseData))
	if err != nil {
		http.Error(w, "JSON 编码失败", http.StatusInternalServerError)
		return
	}
	return
}

func getPaginationParams(r *http.Request) (int, int) {
	pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || pageNum <= 0 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	return pageNum, pageSize
}

func getDishDtos(dishes []model.Dish) ([]model.DishDto, error) {
	var dishDtos []model.DishDto
	for _, dish := range dishes {
		var categoryName string
		if err := global.DB.Table("category").Where("id = ?", dish.CategoryId).Select("name").Scan(&categoryName).Error; err != nil {
			return nil, err
		}
		dishDto := model.DishDto{
			Dish:         dish,
			CategoryName: categoryName,
		}
		dishDtos = append(dishDtos, dishDto)
	}
	return dishDtos, nil
}
