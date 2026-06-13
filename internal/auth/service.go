package auth

import (
	"context"
	"errors"
	"fmt"
	"server/pkg/jwt"
	"time"

	"server/internal/user"
	"server/pkg/wechat"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Service struct {
	userRepo *user.Repo
	wechat   *wechat.Client
	jwt      *jwt.Service
	redis    *redis.Client
}

func NewService(
	userRepo *user.Repo,
	wechat *wechat.Client,
	jwt *jwt.Service,
	redis *redis.Client,
) *Service {

	return &Service{
		userRepo: userRepo,
		wechat:   wechat,
		jwt:      jwt,
		redis:    redis,
	}
}

func (s *Service) Login(
	req LoginRequest,
) (*LoginResponse, error) {

	switch req.Provider {

	case ProviderMP:
		return s.loginMiniProgram(req)

	case ProviderWeb:
		return nil, errors.New("web login not implemented")

	default:
		return nil, errors.New("unsupported provider")
	}
}

func (s *Service) loginMiniProgram(
	req LoginRequest,
) (*LoginResponse, error) {

	session, err := s.wechat.Code2Session(req.Code)

	if err != nil {
		return nil, err
	}

	u, err := s.userRepo.GetByOpenID(
		session.OpenID,
	)

	if errors.Is(
		err,
		gorm.ErrRecordNotFound,
	) {

		u = &user.User{
			OpenID: &session.OpenID,
			Name:   "微信用户",
		}

		if err := s.userRepo.Create(u); err != nil {
			return nil, err
		}

	} else if err != nil {
		return nil, err
	}

	sid := uuid.NewString()
	token, err := s.jwt.Generate(u.ID, sid)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf(
		"session:%d:%s",
		u.ID,
		sid,
	)

	err = s.redis.Set(
		context.Background(),
		key,
		"1",
		30*24*time.Hour,
	).Err()

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
		User: UserResponse{
			ID:     u.ID,
			Name:   u.Name,
			Email:  u.Email,
			Avatar: u.Avatar,
		},
	}, nil
}
