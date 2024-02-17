package user_domain

import (
	"bytes"
	"time"

	"github.com/google/uuid"
)

type MFASetting struct {
	UserUUID  uuid.UUID    `gorm:"column:user_uuid;type:uuid"`
	Key       *string      `gorm:"column:key"`
	CreatedAt *time.Time   `gorm:"column:created_at"`
	KeyImage  bytes.Buffer `gorm:"-:all"`
	// LastLogUUID *uuid.UUID `db:"last_log_uuid"`
}
