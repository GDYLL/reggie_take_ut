package session

import (
	"github.com/gorilla/sessions"
)

var (
	Store       = sessions.NewCookieStore([]byte("session-key"))
	SessionName = "session_id"
)

func init() {
	// 配置全局session选项
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // 1天
		HttpOnly: true,
	}
}
