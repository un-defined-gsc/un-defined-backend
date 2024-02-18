package social_domain

import (
	"time"

	"github.com/google/uuid"
)

// Comment ...
type Comment struct {
	ID        *uuid.UUID `gorm:"column:id,primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `gorm:"cloumn:created_at" json:"created_at"`
	UserID    uuid.UUID  `gorm:"cloumn:user_id" json:"user_id"`
	PostID    uuid.UUID  `gorm:"cloumn:post_id" json:"post_id"`
	Body      string     `gorm:"cloumn:body" json:"body"`
}
