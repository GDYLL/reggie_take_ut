package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/internal/handler"
)

func employeeRouter(r *gin.Engine) {
	emp := r.Group("/employee")
	{
		emp.POST("/login", handler.EmployeeController{}.Login())
		emp.POST("", handler.EmployeeController{}.Save())
		emp.GET("/page", handler.EmployeeController{}.Page())
		emp.GET(":id", handler.EmployeeController{}.Get())
		emp.PUT("", handler.EmployeeController{}.Update())
	}
}
