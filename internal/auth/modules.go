package auth

import (
	"server/internal/user"
	"server/pkg/jwt"
	"server/pkg/wechat"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Module struct {
	Handler *Handler
	Service *Service
}

func NewModule(
	db *gorm.DB,
	wx *wechat.Client,
	jwt *jwt.Service,
	redis *redis.Client,
) *Module {

	userRepo := user.NewRepo(db)

	service := NewService(
		userRepo,
		wx,
		jwt,
		redis,
	)

	handler := NewHandler(
		service,
	)

	return &Module{
		Handler: handler,
		Service: service,
	}
}
