package middleware

import (
	"server/internal/rbac"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...rbac.Role) gin.HandlerFunc {

	return func(c *gin.Context) {

		roleVal, _ := c.Get("role")
		role := roleVal.(rbac.Role)

		for _, r := range roles {
			if role == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
	}
}
