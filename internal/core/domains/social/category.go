package social_domain

import (
	"time"

	"github.com/google/uuid"
)

// Category ...
type Category struct {
	UUID      *uuid.UUID `gorm:"column:uuid,primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	CreatedAt time.Time  `json:"created_at"`
	Name      string     `gorm:"column:name" json:"category"`
}
