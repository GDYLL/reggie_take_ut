package router

import (
	"net/http"
	"reggie_take_ut/internal/handler"
)

func dishRouter(mux *http.ServeMux) {

	mux.HandleFunc("/dish/page", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.DishController{}.Page(w, r)
		}
	})

}
