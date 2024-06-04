package config

import (
	"log"

	"github.com/gin-gonic/gin"
)

// WebMvcConfig 配置静态资源映射
func WebMvcConfig(router *gin.Engine) {
	log.Println("开始进行静态资源映射...")

	// 设置静态资源映射
	router.Static("/backend", "./static/backend")
	router.Static("/front", "./static/front")
}
