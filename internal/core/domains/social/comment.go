package social_domain

import (
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

// Comment ...
type Comment struct {
	domains.Base
	UserID uuid.UUID `gorm:"cloumn:user_id" json:"user_id"`
	PostID uuid.UUID `gorm:"cloumn:post_id" json:"post_id"`
	Body   string    `gorm:"cloumn:body" json:"body"`
}
