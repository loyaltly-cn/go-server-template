package app

import (
	"server/internal/banner"
	"server/internal/user"

	"gorm.io/gorm"
)

type Modules struct {
	User   *user.Module
	Banner *banner.Module
}

func NewModules(db *gorm.DB) *Modules {

	return &Modules{
		User:   user.NewModule(db),
		Banner: banner.NewModule(db),
	}
}
