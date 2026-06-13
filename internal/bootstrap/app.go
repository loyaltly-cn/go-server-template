package bootstrap

import (
	"server/internal/app"
	"server/internal/database"
	"server/pkg/config"
	"server/pkg/id"
	"server/pkg/jwt"
	"server/pkg/wechat"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Config  *config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	Modules *app.Modules
	JWT     *jwt.Service
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

	wxClient := wechat.NewClient(
		config.GetEnv("WECHAT_APPID"),
		config.GetEnv("WECHAT_SECRET"),
	)

	jwtService := jwt.New(config.GetEnv("JWT_SECRET"))

	modules := app.NewModules(db, wxClient, jwtService, rdb)

	return &App{
		Config:  cfg,
		DB:      db,
		Redis:   rdb,
		Modules: modules,
		JWT:     jwtService,
	}
}
