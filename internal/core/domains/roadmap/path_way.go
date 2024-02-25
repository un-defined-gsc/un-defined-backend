package roadmap_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// PathWay ...
type PathWay struct {
	base_domain.Base
	RoadmapID   uuid.UUID `db:"roadmap_id" json:"roadmap_id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	ParentID    uuid.UUID `db:"parent_id" json:"parent_id"`
}
