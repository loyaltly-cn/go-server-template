package middleware

import (
	"net/http"
	"strings"

	"server/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type JWTMiddleware struct {
	jwt *jwt.Service
}

func NewJWTMiddleware(jwt *jwt.Service) *JWTMiddleware {
	return &JWTMiddleware{jwt: jwt}
}

func (m *JWTMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "missing token"})
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid token format"})
			return
		}

		claims, err := m.jwt.Parse(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
