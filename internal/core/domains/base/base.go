package base_domain

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        *uuid.UUID `gorm:"column:id,primaryKey;type:uuid;default:uuid_generate_v4()" json:"-"`
	CreatedAt *time.Time `gorm:"column:created_at;default:now()" json:"-"`
}
