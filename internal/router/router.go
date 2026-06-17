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

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// global middleware
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.JSONOnly())
	r.Use(middleware.APIKeyAuth())
	r.Use(gin.Recovery())

	// health check（不需要 api 前缀）
	r.GET("/ping", func(c *gin.Context) {
		common.Success("pong")
	})

	// =========================
	// PUBLIC API（统一 /api）
	// =========================

	api := r.Group("/api")

	// auth（登录不需要 jwt）
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/login", app.Modules.Auth.Handler.Login)
	}

	// users（公开查询）
	userGroup := api.Group("/users")
	{
		userGroup.GET("/:id", app.Modules.User.Handler.GetUser)
	}

	// banners（公开查询）
	bannerPublic := api.Group("/banners")
	{
		bannerPublic.POST("/query", app.Modules.Banner.Handler.Query)
	}

	// =========================
	// PRIVATE API（需要 JWT）
	// =========================

	private := api.Group("")
	private.Use(middleware.NewJWTMiddleware(app.JWT).Handler())
	{

		authPrivate := private.Group("/auth")
		{
			authPrivate.POST("/auto", app.Modules.Auth.Handler.Me)
		}

		bannerPrivate := private.Group("/banners")
		{
			bannerPrivate.POST("", app.Modules.Banner.Handler.Create)
			bannerPrivate.PUT("/:id", app.Modules.Banner.Handler.Update)
			bannerPrivate.PATCH("/:id", app.Modules.Banner.Handler.Patch)
			bannerPrivate.DELETE("/:id", app.Modules.Banner.Handler.Delete)
		}
	}

	return r
}
