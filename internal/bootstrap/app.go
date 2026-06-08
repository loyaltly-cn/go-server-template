package bootstrap

import (
	"server/internal/app"
	"server/internal/database"
	"server/pkg/config"
	"server/pkg/id"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Config  *config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	Modules *app.Modules
}

func NewApp(cfg *config.Config) *App {

	if err := id.InitSnowflake(1); err != nil {
		panic(err)
	}

	rdb := database.NewRedis(cfg)
	db := database.NewPostgres(cfg)

	if err := database.Migrate(db); err != nil {
		panic(err)
	}

	modules := app.NewModules(db)

	return &App{
		Config:  cfg,
		DB:      db,
		Redis:   rdb,
		Modules: modules,
	}
}
