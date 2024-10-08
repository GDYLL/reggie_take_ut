package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/global"
	"gorm.io/gorm"
	"net/http"
	"reggie_take_ut/internal/model"
	"reggie_take_ut/pkg/common"
	"strconv"
)

type DishController struct {
}

func (dc DishController) Page() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNum, pageSize := getPaginationParams(c)
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}

		dishDtos, err := getDishDtos(dishes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}

		responseData := model.ResponseData{
			Records: dishDtos,
			Total:   total,
		}

		c.JSON(http.StatusOK, common.Success(responseData))
	}
}

func getPaginationParams(c *gin.Context) (int, int) {
	pageNum, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || pageNum <= 0 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("size", "10"))
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
