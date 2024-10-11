package router

import (
	"net/http"
	"reggie_take_ut/internal/handler"
)

// employeeRouter 配置员工相关的路由
func employeeRouter(mux *http.ServeMux) {

	mux.HandleFunc("/employee/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.EmployeeController{}.Login(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.EmployeeController{}.Save(w, r)
		case http.MethodPut:
			handler.EmployeeController{}.Update(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/employee/page", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.EmployeeController{}.Page(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/employee/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.EmployeeController{}.Get(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/employee/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.EmployeeController{}.Logout(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
