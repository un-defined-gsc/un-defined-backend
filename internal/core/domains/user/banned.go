package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type Banned struct {
	base_domain.Base
	UserID    *uuid.UUID `gorm:"column:user_id;type:UUID;NOT NULL"`
	AdminID   *uuid.UUID `gorm:"column:admin_id;type:UUID;NOT NULL"`
	Reason    string     `gorm:"column:reason;type:TEXT;NOT NULL"`
	Permanent bool       `gorm:"column:permanent;type:BOOLEAN;NOT NULL"`
	ExpiresAt *time.Time `gorm:"column:expires_at;type:TIMESTAMP;NOT NULL"`
}

func (Banned) TableName() string {
	return "t_banneds"
}
