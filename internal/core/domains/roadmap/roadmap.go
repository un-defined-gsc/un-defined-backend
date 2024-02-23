package roadmap_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// Roadmap ...
type Roadmap struct {
	base_domain.Base
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	FirstPathID uuid.UUID `db:"first_path_id" json:"first_path_id"`
}
