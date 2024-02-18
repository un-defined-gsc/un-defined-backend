package social_domain

import (
	"time"

	"github.com/google/uuid"
)

// Like ...
type Like struct {
	ID        *uuid.UUID `gorm:"column:id,primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UserID    uuid.UUID  `gorm:"column:user_id" json:"user_id"`
	PostID    uuid.UUID  `gorm:"column:post_id" json:"post_id"`
}
