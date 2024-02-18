package roadmap_domain

import (
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

// Category ...
type SubPath struct {
	domains.Base
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	ParentPath string    `json:"parent_path"`
	PathWayID  uuid.UUID `json:"pathway_id"`
}
