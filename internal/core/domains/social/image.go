package social_domain

import "github.com/google/uuid"

type Image struct {
	UUID      *uuid.UUID `gorm:"column:uuid,primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	PostID    uuid.UUID  `gorm:"column:post_id" json:"post_id"`
	UserID    uuid.UUID  `gorm:"column:user_id" json:"user_id"`
	Path      string     `gorm:"column:path" json:"path"`
	CreatedAt string     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string     `gorm:"updated_at" json:"updated_at"`
}
