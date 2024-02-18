package social_domain

import (
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

type Image struct {
	domains.Base
	PostID uuid.UUID `gorm:"column:post_id" json:"post_id"`
	UserID uuid.UUID `gorm:"column:user_id" json:"user_id"`
	Path   string    `gorm:"column:path" json:"path"`
}
