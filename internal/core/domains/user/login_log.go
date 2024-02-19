package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type LoginLog struct {
	base_domain.Base
	UserID    *uuid.UUID `db:"user_id"`
	LoginAt   *time.Time `db:"login_at"`
	IPAddress string     `db:"ip_address"`
	UserAgent string     `db:"user_agent"`
	IsSuccess bool       `db:"is_success"`
}
