package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/global"
	"net/http"
	"reggie_take_ut/internal/model"
	"reggie_take_ut/pkg/common"
	"strconv"
)

type CategoryController struct{}

func (c CategoryController) Page() gin.HandlerFunc {
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
		var category []model.Category
		if global.DB.Table("category").Offset(offset).Limit(pageSize).Find(&category).Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}
		var total int64
		if global.DB.Table("category").Count(&total).Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}
		responseData := model.ResponseData{
			Records: category,
			Total:   total,
		}
		c.JSON(http.StatusOK, common.Success(responseData))
	}
}
