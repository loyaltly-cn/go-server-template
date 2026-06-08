package banner

import (
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
