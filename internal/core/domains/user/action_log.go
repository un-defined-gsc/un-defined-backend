package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type ActionLog struct {
	UUID       *uuid.UUID `db:"uuid"`
	UserUUID   *uuid.UUID `db:"user_uuid"`
	ActionAt   *time.Time `db:"action_at"`
	ActionSlug string     `db:"action_slug"`
	IPAddress  string     `db:"ip_address"`
	UserAgent  string     `db:"user_agent"`
}
