package social_domain

import (
	"time"

	"github.com/google/uuid"
)

// Category ...
type Category struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"category"`
}
