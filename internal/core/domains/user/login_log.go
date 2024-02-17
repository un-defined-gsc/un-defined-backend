package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type LoginLog struct {
	UUID      *uuid.UUID `gorm:"column:uuid"`
	UserUUID  *uuid.UUID `gorm:"column:user_uuid"`
	LoginAt   *time.Time `gorm:"column:login_at"`
	IPAddress string     `gorm:"column:ip_address"`
	UserAgent string     `gorm:"column:user_agent"`
	IsSuccess bool       `gorm:"column:is_success"`
}
