package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/global"
	"gorm.io/gorm"
	"net/http"
	"reggie_take_ut/common"
	"reggie_take_ut/entity"
	"strconv"
	"time"
)

type CategoryController struct {
}

type ResponseData struct {
	Records []entity.Category `json:"records"`
	Total   int64             `json:"total"`
}

func (e CategoryController) Save() gin.HandlerFunc {
	return func(context *gin.Context) {
		var cgInput entity.Category
		if err := context.ShouldBindJSON(&cgInput); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		var category entity.Category
		if err := global.DB.Table("category").Where("name = ?", cgInput.Name).First(&category).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 设置创建时间和更新时间
				now := time.Now()
				// 填充Employee结构体的字段
				cgInput.CreateTime = now
				cgInput.UpdateTime = now

				if err := global.DB.Table("category").Create(&cgInput).Error; err != nil {
					// 处理错误
					context.JSON(http.StatusInternalServerError, gin.H{
						"error": "Error inserting new category record",
					})
					return
				}
				// 返回成功的响应
				context.JSON(http.StatusOK, common.Success("保存成功"))
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": "Database query failed",
				})
				return
			}
		} else {
			context.JSON(http.StatusOK, common.Error("菜品分类名称已存在"))
			return

		}
	}
}
func (e CategoryController) Page() gin.HandlerFunc {
	return func(context *gin.Context) {
		page := context.DefaultQuery("page", "1")
		pageSize := context.DefaultQuery("pageSize", "10")

		pageNum, err := strconv.Atoi(page)
		if err != nil || pageNum <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "无效的页码"})
			return
		}

		pageSizeNum, err := strconv.Atoi(pageSize)
		if err != nil || pageSizeNum <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "无效的页大小"})
			return
		}

		offset := (pageNum - 1) * pageSizeNum
		var categories []entity.Category
		err = global.DB.Table("category").Offset(offset).Limit(pageSizeNum).Find(&categories).Error
		if err != nil {
			// 如果查询失败，返回错误信息
			context.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}
		var total int64
		err = global.DB.Table("category").Count(&total).Error
		if err != nil {
			// 如果查询失败，返回错误信息
			context.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}
		responseData := ResponseData{
			Records: categories,
			Total:   total,
		}
		context.JSON(http.StatusOK, common.Success(responseData))
	}
}
