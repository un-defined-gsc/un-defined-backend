package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type LoginLog struct {
	UUID      *uuid.UUID `gorm:"column:uuid,primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	UserUUID  *uuid.UUID `gorm:"column:user_uuid;type:uuid"`
	LoginAt   *time.Time `gorm:"column:login_at"`
	IPAddress string     `gorm:"column:ip_address"`
	UserAgent string     `gorm:"column:user_agent"`
	IsSuccess bool       `gorm:"column:is_success"`
}
