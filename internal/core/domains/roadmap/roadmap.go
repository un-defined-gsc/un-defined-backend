package roadmap_domain

import base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"

// Roadmap ...
type Roadmap struct {
	base_domain.Base
	Title string `json:"title"`
	Body  string `json:"body"`
}
