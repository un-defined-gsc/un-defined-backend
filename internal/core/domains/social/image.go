package social_domain

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	ID        uuid.UUID  `db:"id;primary_key" json:"id"`
	PostID    uuid.UUID  `db:"post_id" json:"post_id"`
	UserID    uuid.UUID  `db:"user_id" json:"user_id"`
	Category  string     `db:"category" json:"category"`
	Path      string     `db:"path" json:"path"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}
