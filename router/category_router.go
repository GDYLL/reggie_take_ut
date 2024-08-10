package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/controller"
)

func categoryRouter(r *gin.Engine) {
	cat := r.Group("/category")
	{
		cat.GET("/page", controller.CategoryController{}.Page())
		//cat.POST("", controller.CategoryController{}.Save())
	}
}
