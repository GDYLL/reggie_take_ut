package router

import (
	"github.com/gin-gonic/gin"
	"reggie_take_ut/controller"
)

func dishRouter(r *gin.Engine) {

	dish := r.Group("/dish")
	{
		dish.GET("/page", controller.DishController{}.Page())
	}

}
