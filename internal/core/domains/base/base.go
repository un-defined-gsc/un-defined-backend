package base_domain

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        *uuid.UUID `gorm:"column:id;primaryKey;type:UUID;default:uuid_generate_v4()" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
