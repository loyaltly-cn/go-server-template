package app

import (
	"server/internal/auth"
	"server/internal/banner"
	"server/internal/user"
	"server/pkg/jwt"
	"server/pkg/wechat"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Modules struct {
	User   *user.Module
	Banner *banner.Module
	Auth   *auth.Module
}

func NewModules(
	db *gorm.DB,
	wx *wechat.Client,
	jwtSvc *jwt.Service,
	redis *redis.Client,
) *Modules {

	return &Modules{
		User:   user.NewModule(db),
		Banner: banner.NewModule(db),
		Auth:   auth.NewModule(db, wx, jwtSvc, redis),
	}
}
