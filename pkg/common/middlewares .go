package common

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"path"
	"reggie_take_ut/pkg/session"
	"strings"
)

// 定义不需要处理的请求路径
var excludedURLs = []string{
	"/employee/login",
	"/employee/logout",
	"/backend/**",
	"/front/**",
	"/common/**",
	"/user/sendMsg",
	"/user/login",
	"/doc.html",
	"/webjars/**",
	"/swagger-resources",
	"/v2/api-docs",
}

func JSONError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Result{}.Error(message))
}

func LoginCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("检查请求: %s %s", r.Method, r.URL.Path)

		// 检查是否是排除的路径
		if checkPath(excludedURLs, r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		// 获取session
		sess, err := session.Store.Get(r, session.SessionName)
		if err != nil {
			log.Printf("获取session失败: %v", err)
			JSONError(w, http.StatusInternalServerError, "NOTLOGIN")
			return
		}

		// 检查 session 中的用户ID和登录状态
		employeeID, ok := sess.Values["employee_id"]
		authenticated, authOk := sess.Values["authenticated"].(bool)

		if !ok || !authOk || !authenticated {
			log.Printf("用户未登录或session无效")
			JSONError(w, http.StatusUnauthorized, "NOTLOGIN")
			return
		}

		log.Printf("用户已登录: %v", employeeID)

		// 设置用户信息到上下文
		ctx := context.WithValue(r.Context(), "employee_id", employeeID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// checkPath 检查路径是否匹配排除列表
func checkPath(urls []string, requestURI string) bool {
	for _, url := range urls {
		matched, _ := path.Match(strings.Replace(url, "**", "*", -1), requestURI)
		if matched {
			return true
		}
	}
	return false
}
