package router

import (
	"net/http"
	"reggie_take_ut/internal/handler"
)

func commonRouter(mux *http.ServeMux) {
	//common := r.Group("/common")
	//{
	//	common.GET("/download", handler.CommonController{}.Download)
	//}
	mux.HandleFunc("/common/download", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.CommonController{}.Download(w, r)
		}
	})
}
