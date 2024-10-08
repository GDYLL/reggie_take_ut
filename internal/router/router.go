package router

import (
	"github.com/spf13/viper"
	"net/http"
	"reggie_take_ut/config"
)

// RouteGroup 用于组织多个路由
type RouteGroup struct {
	prefix string
	mux    *http.ServeMux
}

// NewRouteGroup 创建新的路由组
func NewRouteGroup(mux *http.ServeMux, prefix string) *RouteGroup {
	return &RouteGroup{prefix: prefix, mux: mux}
}

// Handle 用于注册 GET/POST 等 HTTP 方法的路由
func (g *RouteGroup) Handle(path string, handler http.HandlerFunc) {
	fullPath := g.prefix + path
	g.mux.HandleFunc(fullPath, handler)
}

func InitRouter() {

	mux := http.NewServeMux()

	config.WebMvcConfig(mux)

	categoryGroup := NewRouteGroup(mux, "/category")
	categoryRouter(categoryGroup)

	employeeRouter(mux)

	dishRouter(mux)

	commonRouter(mux)

	port := viper.GetString("server.port")
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		return
	}
}
