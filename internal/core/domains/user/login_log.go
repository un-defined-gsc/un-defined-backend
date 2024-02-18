package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type LoginLog struct {
	base_domain.Base
	UserID    *uuid.UUID `gorm:"column:user_id;type:UUID;NOT NULL"`
	LoginAt   *time.Time `gorm:"column:login_at;type:TIMESTAMP;NOT NULL;default:CURRENT_TIMESTAMP"`
	IPAddress string     `gorm:"column:ip_address;type:TEXT;NOT NULL"`
	UserAgent string     `gorm:"column:user_agent;type:TEXT;NOT NULL"`
	IsSuccess bool       `gorm:"column:is_success;type:BOOLEAN;NOT NULL"`
}

func (LoginLog) TableName() string {
	return "t_login_logs"
}
