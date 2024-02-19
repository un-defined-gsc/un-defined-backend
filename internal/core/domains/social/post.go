package social_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// Post is our main model for Posts
type Post struct {
	base_domain.Base
	CategoryID uuid.UUID `json:"category_id"`
	UserID     uuid.UUID `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"body"`
	Tags       []string  `json:"tags"`
}
