package common

import (
	"github.com/gin-gonic/gin"
)

//func SetCookie(CookieName string, Value string, Expires time.Time, c *gin.Context) {
//	cookie := &http.Cookie{
//		Name:     CookieName,
//		Value:    Value,
//		Expires:  Expires,
//		HttpOnly: true,
//	}
//	c.SetCookie(c.Writer, cookie)
//}

func SetCookie(CookieName string, Value string, Expires int, c *gin.Context) {
	// 使用 Gin 的 SetCookie 方法设置 cookie
	c.SetCookie(CookieName, Value, Expires, "/", "", false, true)
}
