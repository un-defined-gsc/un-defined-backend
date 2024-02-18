package roadmap_domain

import (
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

// PathWay ...
type PathWay struct {
	domains.Base
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	RoadmapID  uuid.UUID `json:"roadmap_id"`
	ParentPath uuid.UUID `json:"parent_path"`
}
