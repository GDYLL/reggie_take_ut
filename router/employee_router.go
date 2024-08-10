package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/controller"
)

func employeeRouter(r *gin.Engine) {
	emp := r.Group("/employee")
	{
		emp.POST("/login", controller.EmployeeController{}.Login())
		emp.POST("", controller.EmployeeController{}.Save())
		emp.GET("/page", controller.EmployeeController{}.Page())
		emp.GET(":id", controller.EmployeeController{}.Get())
		emp.PUT("", controller.EmployeeController{}.Update())
	}
}
