package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type ActionLog struct {
	ID         *uuid.UUID `gorm:"column:id,primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID     *uuid.UUID `gorm:"column:user_id;type:uuid"`
	ActionAt   *time.Time `gorm:"column:action_at"`
	ActionSlug string     `gorm:"column:action_slug"`
	IPAddress  string     `gorm:"column:ip_address"`
	UserAgent  string     `gorm:"column:user_agent"`
}
