package router

import (
	"net/http"
	"reggie_take_ut/internal/handler"
)

func categoryRouter(mux *http.ServeMux) {
	mux.HandleFunc("/category", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CategoryController{}.Save(w, r)
		case http.MethodPut:
			handler.CategoryController{}.Update(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/category/page", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.CategoryController{}.Page(w, r)
		}
	})
}
