package router

import (
	"github.com/spf13/viper"
	"net/http"
	"reggie_take_ut/config"
	"reggie_take_ut/pkg/common"
)

// RouteGroup 用于组织多个路由
type RouteGroup struct {
	prefix      string
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

// NewRouteGroup 创建新的路由组
func NewRouteGroup(mux *http.ServeMux, prefix string) *RouteGroup {
	return &RouteGroup{
		prefix:      prefix,
		mux:         mux,
		middlewares: make([]func(http.Handler) http.Handler, 0),
	}
}

// Use 添加中间件
func (g *RouteGroup) Use(middleware ...func(http.Handler) http.Handler) {
	g.middlewares = append(g.middlewares, middleware...)
}

// Handle 用于注册 GET/POST 等 HTTP 方法的路由
func (g *RouteGroup) Handle(path string, handler http.HandlerFunc) {
	fullPath := g.prefix + path

	// 将所有中间件应用到处理器
	var finalHandler http.Handler = handler
	for i := len(g.middlewares) - 1; i >= 0; i-- {
		finalHandler = g.middlewares[i](finalHandler)
	}

	g.mux.Handle(fullPath, finalHandler)
}

// 创建一个包装 ServeMux 的结构体
type Router struct {
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

// 创建新的 Router
func NewRouter() *Router {
	return &Router{
		mux:         http.NewServeMux(),
		middlewares: make([]func(http.Handler) http.Handler, 0),
	}
}

// Use 为整个路由器添加中间件
func (r *Router) Use(middleware ...func(http.Handler) http.Handler) {
	r.middlewares = append(r.middlewares, middleware...)
}

// ServeHTTP 实现 http.Handler 接口
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.mux
	// 按照反序应用中间件
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}
	handler.ServeHTTP(w, req)
}

func InitRouter() {
	// 创建新的路由器
	router := NewRouter()
	mux := router.mux

	router.Use(common.DebugMiddleware)

	// 配置 Web MVC
	config.WebMvcConfig(mux)

	categoryRouter(mux)
	employeeRouter(mux)
	dishRouter(mux)
	commonRouter(mux)

	// 添加全局登录检查中间件
	router.Use(common.LoginCheckMiddleware)

	// 启动服务器
	port := viper.GetString("server.port")
	err := http.ListenAndServe(":"+port, router) // 使用包装后的 router
	if err != nil {
		return
	}
}
