package router

import (
	"reggie_take_ut/internal/handler"
)

func categoryRouter(r *RouteGroup) {
	r.Handle("/page", handler.CategoryController{}.Page)
}
