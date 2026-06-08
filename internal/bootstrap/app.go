package bootstrap

import (
	"server/internal/database"
	"server/internal/user"
	"server/pkg/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config

	DB    *gorm.DB
	Redis *redis.Client
	User  *user.Module
}

func NewApp(cfg *config.Config) *App {

	db := database.NewPostgres(cfg)
	rdb := database.NewRedis(cfg)

	return &App{
		Config: cfg,
		DB:     db,
		Redis:  rdb,
	}
}
