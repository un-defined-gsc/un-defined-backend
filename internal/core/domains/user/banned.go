package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type Banned struct {
	UUID      *uuid.UUID `gorm:"column:uuid,primaryKey;type:uuid" json:"uuid"`
	UserUUID  *uuid.UUID `gorm:"column:user_uuid;type:uuid"`
	AdminUUID *uuid.UUID `gorm:"column:admin_uuid;type:uuid"`
	Reason    string     `gorm:"column:reason"`
	Permanent bool       `gorm:"column:permanent"`
	ExpiresAt *time.Time `gorm:"column:expires_at"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}
