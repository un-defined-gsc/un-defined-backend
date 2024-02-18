package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type ActionLog struct {
	base_domain.Base
	UserID     *uuid.UUID `gorm:"column:user_id;type:UUID;NOT NULL"`
	ActionAt   *time.Time `gorm:"column:action_at;type:TIMESTAMP;NOT NULL;default:CURRENT_TIMESTAMP"`
	ActionSlug string     `gorm:"column:action_slug;type:TEXT;NOT NULL"`
	IPAddress  string     `gorm:"column:ip_address;type:TEXT;NOT NULL"`
	UserAgent  string     `gorm:"column:user_agent;type:TEXT;NOT NULL"`
}

func (ActionLog) TableName() string {
	return "t_action_logs"
}
