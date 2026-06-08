package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JSONOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		method := c.Request.Method

		// 只检查有 body 的请求
		if method == "POST" || method == "PUT" || method == "PATCH" {

			contentType := c.GetHeader("Content-Type")

			if !strings.HasPrefix(contentType, "application/json") {
				c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
					"error": "content-type must be application/json",
				})
				return
			}
		}

		c.Next()
	}
}
