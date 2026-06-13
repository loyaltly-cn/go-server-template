package jwt

import (
	"time"

	gjwt "github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secret []byte
}

func New(secret string) *Service {
	return &Service{secret: []byte(secret)}
}

type Claims struct {
	UserID    int64  `json:"user_id"`
	SessionID string `json:"sid"`
	gjwt.RegisteredClaims
}

func (s *Service) Generate(userID int64, sid string) (string, error) {

	claims := Claims{
		UserID:    userID,
		SessionID: sid,
		RegisteredClaims: gjwt.RegisteredClaims{
			ExpiresAt: gjwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  gjwt.NewNumericDate(time.Now()),
		},
	}

	token := gjwt.NewWithClaims(gjwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *Service) Parse(tokenStr string) (*Claims, error) {

	token, err := gjwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(t *gjwt.Token) (interface{}, error) {
			return s.secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
