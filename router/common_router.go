package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/controller"
)

func commonRouter(r *gin.Engine) {
	common := r.Group("/common")
	{
		common.GET("/download", controller.CommonController{}.Download)
	}
}
