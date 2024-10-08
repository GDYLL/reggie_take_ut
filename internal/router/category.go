package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/internal/handler"
)

func categoryRouter(r *gin.Engine) {
	cat := r.Group("/category")
	{
		cat.GET("/page", handler.CategoryController{}.Page())
		//cat.POST("", controller.CategoryController{}.Save())
	}
}
