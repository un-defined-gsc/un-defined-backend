package social_domain

import (
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

// Like ...
type Like struct {
	domains.Base
	UserID uuid.UUID `gorm:"column:user_id" json:"user_id"`
	PostID uuid.UUID `gorm:"column:post_id" json:"post_id"`
}
