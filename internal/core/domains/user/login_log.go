package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type LoginLog struct {
	base_domain.Base
	UserID    *uuid.UUID `gorm:"column:user_id;type:uuid"`
	LoginAt   *time.Time `gorm:"column:login_at"`
	IPAddress string     `gorm:"column:ip_address"`
	UserAgent string     `gorm:"column:user_agent"`
	IsSuccess bool       `gorm:"column:is_success"`
}
