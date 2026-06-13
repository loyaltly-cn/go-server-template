package user

import (
	"server/internal/rbac"
	"server/pkg/id"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID int64 `gorm:"primaryKey"`

	Name     string
	Email    *string `gorm:"uniqueIndex"`
	Password string
	OpenID   *string `gorm:"uniqueIndex"`
	Phone    *string `gorm:"uniqueIndex"`
	Avatar   string  `gorm:"column:avatar"`

	Role rbac.Role `gorm:"type:varchar(20);default:'user'"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == 0 {
		u.ID = id.GenerateID()
	}
	return nil
}
