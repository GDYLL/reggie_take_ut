package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/controller"
)

func categoryRouter(r *gin.Engine) {
	cg := r.Group("/category")
	{
		cg.POST("", controller.CategoryController{}.Save())
		cg.GET("/page", controller.CategoryController{}.Page())
	}
}
