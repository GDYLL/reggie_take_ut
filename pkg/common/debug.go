package common

import (
	"log"
	"net/http"
	"reggie_take_ut/pkg/session"
)

func DebugMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("=== 请求开始 ===")
		log.Printf("方法: %s", r.Method)
		log.Printf("路径: %s", r.URL.Path)
		log.Printf("Cookie: %v", r.Cookies())

		// 尝试获取session
		if sess, err := session.Store.Get(r, "employee-session"); err == nil {
			log.Printf("Session值: %v", sess.Values)
		} else {
			log.Printf("获取session失败: %v", err)
		}

		next.ServeHTTP(w, r)
		log.Printf("=== 请求结束 ===\n")
	})
}
