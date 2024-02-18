package social_domain

import (
	"time"

	"github.com/google/uuid"
)

// Post is our main model for Posts
type Post struct {
	ID         *uuid.UUID `gorm:"column:id,primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CategoryID uuid.UUID  `json:"category_id"`
	UserID     uuid.UUID  `json:"user_id"`
	CreatedAt  time.Time  `json:"created_at"`
	Title      string     `json:"title"`
	Body       string     `json:"body"`
}
