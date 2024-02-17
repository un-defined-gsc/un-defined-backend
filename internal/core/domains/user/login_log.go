package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type LoginLog struct {
	UUID      *uuid.UUID `db:"uuid"`
	UserUUID  *uuid.UUID `db:"user_uuid"`
	LoginAt   *time.Time `db:"login_at"`
	IPAddress string     `db:"ip_address"`
	UserAgent string     `db:"user_agent"`
	IsSuccess bool       `db:"is_success"`
}
