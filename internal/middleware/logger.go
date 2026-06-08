package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		cost := time.Since(start)

		rid, _ := c.Get(RequestIDKey)

		log.Printf("[REQ] id=%v method=%s path=%s cost=%v",
			rid,
			c.Request.Method,
			c.Request.URL.Path,
			cost,
		)
	}
}
