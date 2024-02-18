package roadmap_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// Category ...
type SubPath struct {
	base_domain.Base
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	ParentPath string    `json:"parent_path"`
	PathWayID  uuid.UUID `json:"pathway_id"`
}
