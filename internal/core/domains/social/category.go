package social_domain

import (
	"time"

	"github.com/google/uuid"
)

// Category ...
type Category struct {
	ID        *uuid.UUID `gorm:"column:id,primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Name      string     `gorm:"column:name" json:"category"`
}
