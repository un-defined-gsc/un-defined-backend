package social_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// Post is our main model for Posts
type Post struct {
	base_domain.Base
	CategoryID uuid.UUID `db:"category_id"json:"category_id"`
	UserID     uuid.UUID `db:"user_id"json:"user_id"`
	Title      string    `db:"title"json:"title"`
	Content    string    `db:"content"json:"body"`
}
