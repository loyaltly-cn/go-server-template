package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {

		rid := c.GetHeader(RequestIDKey)

		if rid == "" {
			rid = uuid.New().String()
		}

		// 存入 context
		c.Set(RequestIDKey, rid)

		// 返回给客户端
		c.Writer.Header().Set(RequestIDKey, rid)

		c.Next()
	}
}
