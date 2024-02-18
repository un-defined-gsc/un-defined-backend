package roadmap_domain

import (
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

// Roadmap ...
type Roadmap struct {
	domains.Base
	Title string `json:"title"`
	Body  string `json:"body"`
}
