package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type ActionLog struct {
	UUID       *uuid.UUID `gorm:"column:uuid,primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	UserUUID   *uuid.UUID `gorm:"column:user_uuid;type:uuid"`
	ActionAt   *time.Time `gorm:"column:action_at"`
	ActionSlug string     `gorm:"column:action_slug"`
	IPAddress  string     `gorm:"column:ip_address"`
	UserAgent  string     `gorm:"column:user_agent"`
}
