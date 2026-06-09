package banner

import (
	"server/pkg/id"
	"time"

	"gorm.io/gorm"
)

type Banner struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Image     string         `gorm:"not null" json:"image"` // 图片URL
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *Banner) BeforeCreate(tx *gorm.DB) error {
	if u.ID == 0 {
		u.ID = id.GenerateID()
	}
	return nil
}
