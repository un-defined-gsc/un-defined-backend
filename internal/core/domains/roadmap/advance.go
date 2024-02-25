package roadmap_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type Advance struct {
	base_domain.Base
	RoadmapID   uuid.UUID `db:"roadmap_id" json:"roadmap_id"`
	PathWayID   uuid.UUID `db:"pathway_id" json:"pathway_id"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	AdvanceType string    `db:"advance_type" json:"advance_type"`
}
