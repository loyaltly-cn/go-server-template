package jwt

import (
	"server/internal/rbac"
)

type Claims struct {
	UserID int64     `json:"user_id"`
	Role   rbac.Role `json:"role"`
	Source string    `json:"source"`
}
