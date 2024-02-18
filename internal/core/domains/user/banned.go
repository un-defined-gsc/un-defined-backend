package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type Banned struct {
	base_domain.Base
	UserID    *uuid.UUID `gorm:"column:user_id;type:uuid"`
	AdminID   *uuid.UUID `gorm:"column:admin_id;type:uuid"`
	Reason    string     `gorm:"column:reason"`
	Permanent bool       `gorm:"column:permanent"`
	ExpiresAt *time.Time `gorm:"column:expires_at"`
}
