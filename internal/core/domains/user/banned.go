package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type Banned struct {
	ID        *uuid.UUID `gorm:"column:id,primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID    *uuid.UUID `gorm:"column:user_id;type:uuid"`
	AdminID   *uuid.UUID `gorm:"column:admin_id;type:uuid"`
	Reason    string     `gorm:"column:reason"`
	Permanent bool       `gorm:"column:permanent"`
	ExpiresAt *time.Time `gorm:"column:expires_at"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}
