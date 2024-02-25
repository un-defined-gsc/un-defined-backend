package roadmap_domain

import (
	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

// Roadmap ...
type Roadmap struct {
	base_domain.Base
	Name        string    `db:"name" json:"name" validate:"required"`
	CategoryID  uuid.UUID `db:"category" json:"category" validate:"required,uuid"`
	Description string    `db:"description" json:"description" validate:"required"`
	FirstPathID uuid.UUID `db:"first_path_id" json:"first_path_id" validate:"required,uuid"`
}
