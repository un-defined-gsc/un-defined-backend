package social_domain

import (
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

// Category ...
type Category struct {
	domains.Base
	Name string `gorm:"column:name" json:"category"`
}
