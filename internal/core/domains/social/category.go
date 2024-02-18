package social_domain

import base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"

// Category ...
type Category struct {
	base_domain.Base
	Name string `gorm:"column:name" json:"category"`
}
