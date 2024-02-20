package social_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type Image struct {
	base_domain.Base
	PostID   uuid.UUID `gorm:"column:post_id" json:"post_id"`
	UserID   uuid.UUID `gorm:"column:user_id" json:"user_id"`
	Category string    `gorm:"column:category" json:"category"`
	Path     string    `gorm:"column:path" json:"path"`
}
