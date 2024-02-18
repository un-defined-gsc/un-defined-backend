package social_domain

import (
	"time"

	"github.com/google/uuid"
)

// Like ...
type Like struct {
	UUID      *uuid.UUID `gorm:"column:uuid,primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UserID    uuid.UUID  `gorm:"column:user_id" json:"user_id"`
	PostID    uuid.UUID  `gorm:"column:post_id" json:"post_id"`
}

// LikeDTO ...
type LikeDTO struct {
	UUID     uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `json:"username"`
}
