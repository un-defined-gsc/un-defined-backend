package roadmap_domain

import (
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

// CompletedMap ...
type CompletedMap struct {
	domains.Base
	UserID    uuid.UUID `json:"user_id"`
	PathWayID uuid.UUID `json:"pathway_id"`
	RoadmapID uuid.UUID `json:"roadmap_id"`
	SubPathID uuid.UUID `json:"subpath_id"`
}
