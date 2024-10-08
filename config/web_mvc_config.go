package config

import (
	"log"
	"net/http"
)

// WebMvcConfig 配置静态资源映射
func WebMvcConfig(mux *http.ServeMux) {
	log.Println("开始进行静态资源映射...")

	// 设置静态资源映射
	mux.Handle("/backend/", http.StripPrefix("/backend/", http.FileServer(http.Dir("./static/backend"))))
	mux.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("./static/front"))))
}
