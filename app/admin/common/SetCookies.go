package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetCookie(uname string, c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "userName",
		Value:    uname,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Cookie已设置",
	})
}
