package controller

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/global"
	"github.com/prynnekey/go-reggie/utils"
	"reggie_take_ut/common"
	"reggie_take_ut/entity"
)

type EmployeeController struct {
}

func (e EmployeeController) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var empInput entity.Employee
		if err := context.ShouldBindJSON(&empInput); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		username := empInput.Username
		password := utils.MD5(empInput.Password)

		var empStored entity.Employee
		if err := global.DB.Table("employee").Where("username = ?", username).First(&empStored).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database query failed",
			})
			return
		}
		if empStored.Status == 0 {
			context.JSON(http.StatusOK, common.Error("该账户未启用，请联系管理员启用"))
			return
		}
		if empStored.Password == password {
			context.JSON(http.StatusOK, common.Success("登陆成功"))
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
		}
	}
}

func (e EmployeeController) Save() gin.HandlerFunc {
	return func(context *gin.Context) {
		var empInput entity.Employee
		if err := context.ShouldBindJSON(&empInput); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		var empStored entity.Employee
		if err := global.DB.Table("employee").Where("username = ?", empInput.Username).First(&empStored).Error; err != nil {
			// 如果根据username找到记录，则说明已有同名用户，此时应该报错
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 设置默认密码
				password := utils.MD5("123456")
				// 设置创建时间和更新时间
				now := time.Now()
				// 填充Employee结构体的字段
				empInput.Password = password
				empInput.CreateTime = now
				empInput.UpdateTime = now

				if err := global.DB.Table("employee").Create(&empInput).Error; err != nil {
					// 处理错误
					context.JSON(http.StatusInternalServerError, gin.H{
						"error": "Error inserting new employee record",
					})
					return
				}
				// 返回成功的响应
				context.JSON(http.StatusOK, common.Success("保存成功"))
			} else {
				// 查询时出现错误，但不是因为记录未找到
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": "Database query failed",
				})
				return
			}
		} else {
			context.JSON(http.StatusOK, common.Error("用户名已存在"))
			return
		}
	}
}

func (e EmployeeController) Page() gin.HandlerFunc {
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
		var employees []entity.Employee
		err = global.DB.Table("employee").Offset(offset).Limit(pageSizeNum).Find(&employees).Error
		if err != nil {
			// 如果查询失败，返回错误信息
			context.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}
		var total int64
		err = global.DB.Table("employee").Count(&total).Error
		if err != nil {
			// 如果查询失败，返回错误信息
			context.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
			return
		}
		responseData := entity.ResponseData{
			Records: employees,
			Total:   total,
		}
		context.JSON(http.StatusOK, common.Success(responseData))
	}
}
