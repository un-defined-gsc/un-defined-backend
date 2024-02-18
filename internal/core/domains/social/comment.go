package social_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// Comment ...
type Comment struct {
	base_domain.Base
	UserID uuid.UUID `gorm:"cloumn:user_id" json:"user_id"`
	PostID uuid.UUID `gorm:"cloumn:post_id" json:"post_id"`
	Body   string    `gorm:"cloumn:body" json:"body"`
}
