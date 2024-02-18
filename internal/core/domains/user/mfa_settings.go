package user_domain

import (
	"bytes"
	"time"

	"github.com/google/uuid"
)

type MFASetting struct {
	UserID    uuid.UUID    `gorm:"column:user_id;type:UUID;NOT NULL"`
	Key       *string      `gorm:"column:key;type:TEXT;NOT NULL"`
	CreatedAt *time.Time   `gorm:"column:created_at;type:TIMESTAMP;NOT NULL;default:CURRENT_TIMESTAMP"`
	KeyImage  bytes.Buffer `gorm:"-:all"`
	// LastLogUUID *uuid.UUID `db:"last_log_uuid"`
}

func (MFASetting) TableName() string {
	return "t_mfa_settings"
}
