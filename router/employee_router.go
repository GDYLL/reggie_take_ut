package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/controller"
)

func employeeRouter(r *gin.Engine) {
	emp := r.Group("/employee")
	{
		// login
		emp.POST("/login", controller.EmployeeController{}.Login())
		emp.POST("", controller.EmployeeController{}.Save())
	}
}
