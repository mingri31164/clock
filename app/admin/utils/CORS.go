package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许访问的来源，*表示所有来、可以换成具体url
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的 HTTP 方法，*表示所有方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		// 设置允许的 HTTP 头，*表示所有头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// 允许发送 Cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 设置预检请求的缓存时间（秒）
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		// 如果是 OPTIONS 请求，则终止并返回状态码 204(OPTIONS请求为测试请求，不用响应数据，直接在写入响应头之后无内容返回就好了)
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
		} else {
			// 继续处理请求
			c.Next()
		}
	}
}
