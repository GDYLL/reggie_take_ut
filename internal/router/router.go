package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"reggie_take_ut/config"
)

func InitRouter() {

	port := viper.GetString("server.port")
	r := gin.Default()

	// 加载 WebMvcConfig
	config.WebMvcConfig(r)

	employeeRouter(r)
	categoryRouter(r)
	dishRouter(r)
	commonRouter(r)

	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
