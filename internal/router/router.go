package router

import (
	"server/internal/bootstrap"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *bootstrap.App) *gin.Engine {

	r := gin.New()

	// 🔥 1. 全局 middleware（所有请求都会走）
	r.Use(middleware.Logger())
	r.Use(middleware.JSONOnly())
	r.Use(middleware.APIKeyAuth())
	r.Use(gin.Recovery())

	// routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return r
}
