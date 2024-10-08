package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/internal/handler"
)

func dishRouter(r *gin.Engine) {

	dish := r.Group("/dish")
	{
		dish.GET("/page", handler.DishController{}.Page())
	}

}
