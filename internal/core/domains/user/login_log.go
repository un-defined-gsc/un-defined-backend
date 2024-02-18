package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type LoginLog struct {
	ID        *uuid.UUID `gorm:"column:id,primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID    *uuid.UUID `gorm:"column:user_id;type:uuid"`
	LoginAt   *time.Time `gorm:"column:login_at"`
	IPAddress string     `gorm:"column:ip_address"`
	UserAgent string     `gorm:"column:user_agent"`
	IsSuccess bool       `gorm:"column:is_success"`
}
