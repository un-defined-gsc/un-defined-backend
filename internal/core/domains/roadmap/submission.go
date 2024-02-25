package roadmap_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// CompletedMap ...
type Submission struct {
	base_domain.Base
	RoadmapID uuid.UUID `db:"roadmap_id" json:"-"`
	UserID    uuid.UUID `db:"user_id" json:"-"`
}
