package roadmap_domain

import (
	"time"

	"github.com/google/uuid"
)

// Category ...
type SubPath struct {
	ID         uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	ParentPath string    `json:"parent_path"`
	PathWayID  uuid.UUID `json:"pathway_id"`
}

// CategoryDTO ...
type SubPathDTO struct {
	ID    uuid.UUID `gorm:"primary_key" json:"id"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
}

// ToCategory ...
func ToSubPath(subPathDTO *SubPathDTO) *SubPath {
	return &SubPath{
		Title: subPathDTO.Title,
		Body:  subPathDTO.Body,
	}
}

// ToCategoryDTO ...
func ToSubPathDTO(subPath *SubPath) *SubPathDTO {
	return &SubPathDTO{
		ID:    subPath.ID,
		Title: subPath.Title,
		Body:  subPath.Body,
	}
}
