package social_domain

import (
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

// Post is our main model for Posts
type Post struct {
	domains.Base
	CategoryID uuid.UUID `json:"category_id"`
	UserID     uuid.UUID `json:"user_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
}
