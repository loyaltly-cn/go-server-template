package database

import (
	"server/internal/banner"
	"server/internal/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&banner.Banner{},
	)
}
