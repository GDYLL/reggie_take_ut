package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/internal/handler"
)

func commonRouter(r *gin.Engine) {
	common := r.Group("/common")
	{
		common.GET("/download", handler.CommonController{}.Download)
	}
}
