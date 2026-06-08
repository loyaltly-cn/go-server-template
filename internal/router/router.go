package router

import (
	_ "server/docs"
	"server/internal/bootstrap"
	common "server/internal/common/response"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(app *bootstrap.App) *gin.Engine {

	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.JSONOnly())
	r.Use(middleware.APIKeyAuth())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		common.Success("pong")
	})

	userGroup := r.Group("/users")
	{
		userGroup.POST("", app.Modules.User.Handler.CreateUser)
		userGroup.GET("/:id", app.Modules.User.Handler.GetUser)
	}

	bannerGroup := r.Group("/banners")
	{
		bannerGroup.POST("", app.Modules.Banner.Handler.Create)
		bannerGroup.GET("", app.Modules.Banner.Handler.List)
		bannerGroup.GET("/:id", app.Modules.Banner.Handler.Get)
		bannerGroup.PUT("/:id", app.Modules.Banner.Handler.Update)
		bannerGroup.DELETE("/:id", app.Modules.Banner.Handler.Delete)
	}

	return r
}
