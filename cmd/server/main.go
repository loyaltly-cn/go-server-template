// @title Product Support API
// @version 1.0
// @description Product Support Backend API
// @BasePath /
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
package main

import (
	"fmt"
	"server/internal/bootstrap"
	"server/internal/router"
	"server/pkg/config"
)

func main() {

	config.LoadEnv()

	cfg := config.Load()

	app := bootstrap.NewApp(cfg)

	r := router.SetupRouter(app)

	r.Run(fmt.Sprintf(":%d", app.Config.Server.Port))
}
