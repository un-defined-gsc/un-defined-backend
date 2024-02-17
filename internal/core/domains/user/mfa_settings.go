package user_domain

import (
	"bytes"
	"time"

	"github.com/google/uuid"
)

type MFASetting struct {
	UserUUID  uuid.UUID    `db:"user_uuid"`
	Key       *string      `db:"key"`
	CreatedAt *time.Time   `db:"created_at"`
	KeyImage  bytes.Buffer `db:"-"`
	// LastLogUUID *uuid.UUID `db:"last_log_uuid"`
}
