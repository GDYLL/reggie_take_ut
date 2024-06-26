package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/global"
	"net/http"
	"reggie_take_ut/common"
	"reggie_take_ut/entity"
	"strconv"
)

type DishController struct {
}

func (c DishController) Page() gin.HandlerFunc {
	return func(c *gin.Context) {
		num := c.DefaultQuery("page", "1")
		size := c.DefaultQuery("size", "10")

		pageNum, _ := strconv.Atoi(num)
		if pageNum <= 0 {
			pageNum = 1
		}

		pageSize, _ := strconv.Atoi(size)
		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (pageNum - 1) * pageSize
		var dish []entity.Dish
		if global.DB.Table("dish").Offset(offset).Limit(pageSize).Find(&dish).Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}
		var total int64
		if global.DB.Table("dish").Count(&total).Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}

		responseData := entity.ResponseData{
			Records: dish,
			Total:   total,
		}

		c.JSON(http.StatusOK, common.Success(responseData))
	}
}
