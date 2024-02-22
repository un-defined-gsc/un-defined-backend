package social_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// Like ...
type Like struct {
	base_domain.Base
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	PostID uuid.UUID `db:"post_id" json:"post_id"`
}
