package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		rid, _ := c.Get("request_id")

		c.Next()

		latency := time.Since(start)

		status := c.Writer.Status()

		log.Printf(
			"[RID:%v] %s %s %d %v",
			rid,
			c.Request.Method,
			c.Request.URL.Path,
			status,
			latency,
		)
	}
}
