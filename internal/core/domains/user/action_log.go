package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type ActionLog struct {
	base_domain.Base
	UserID     *uuid.UUID `gorm:"column:user_id;type:uuid"`
	ActionAt   *time.Time `gorm:"column:action_at"`
	ActionSlug string     `gorm:"column:action_slug"`
	IPAddress  string     `gorm:"column:ip_address"`
	UserAgent  string     `gorm:"column:user_agent"`
}

func (ActionLog) TableName() string {
	return "t_action_logs"
}
