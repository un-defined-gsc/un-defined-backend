package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type Banned struct {
	UUID      *uuid.UUID `db:"uuid"`
	UserUUID  uuid.UUID  `db:"user_uuid"`
	AdminUUID *uuid.UUID `db:"admin_uuid"`
	Reason    string     `db:"reason"`
	Permanent bool       `db:"permanent"`
	ExpiresAt *time.Time `db:"expires_at"`
	CreatedAt *time.Time `db:"created_at"`
}
