package controller

import (
	"net/http"

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

		if empStored.Password == password {
			context.JSON(http.StatusOK, common.Success("登陆成功"))
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
		}
	}
}
