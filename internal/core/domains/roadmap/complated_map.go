package roadmap_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// CompletedMap ...
type CompletedMap struct {
	base_domain.Base
	UserID    uuid.UUID `json:"user_id"`
	PathWayID uuid.UUID `json:"pathway_id"`
	RoadmapID uuid.UUID `json:"roadmap_id"`
	SubPathID uuid.UUID `json:"subpath_id"`
}
