package user_domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

type ActionLog struct {
	domains.Base
	UserID     *uuid.UUID `gorm:"column:user_id;type:uuid"`
	ActionAt   *time.Time `gorm:"column:action_at"`
	ActionSlug string     `gorm:"column:action_slug"`
	IPAddress  string     `gorm:"column:ip_address"`
	UserAgent  string     `gorm:"column:user_agent"`
}
