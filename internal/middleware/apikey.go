package middleware

import (
	"net/http"
	"server/pkg/config"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing api key",
			})
			return
		}

		if apiKey != config.GetEnv("X_API_KEY") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid api key",
			})
			return
		}

		// 通过校验，继续往下走
		c.Next()
	}
}
