package roadmap_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// PathWay ...
type PathWay struct {
	base_domain.Base
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	RoadmapID  uuid.UUID `json:"roadmap_id"`
	ParentPath uuid.UUID `json:"parent_path"`
}
