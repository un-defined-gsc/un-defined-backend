package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type Banned struct {
	base_domain.Base
	UserID    *uuid.UUID `db:"user_id"`
	AdminID   *uuid.UUID `db:"admin_id"`
	Reason    string     `db:"reason"`
	Permanent bool       `db:"permanent"`
	ExpiresAt *time.Time `db:"expires_at"`
}
