package user_domain

import (
	"time"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type ActionLog struct {
	base_domain.Base
	UserID     *uuid.UUID `db:"user_id"`
	ActionAt   *time.Time `db:"action_at"`
	ActionSlug string     `db:"action_slug"`
	IPAddress  string     `db:"ip_address"`
	UserAgent  string     `db:"user_agent"`
}
