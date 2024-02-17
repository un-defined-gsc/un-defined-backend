package roadmap_domain

import (
	"time"

	"github.com/google/uuid"
)

// PathWay ...
type PathWay struct {
	ID         uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	RoadmapID  uuid.UUID `json:"roadmap_id"`
	ParentPath uuid.UUID `json:"parent_path"`
	SubPath    []SubPath `gorm:"foreignKey:ParentPath;references:ID" json:"sub_path"`
}

// PathWayDTO ...
type PathWayDTO struct {
	ID    uuid.UUID `gorm:"primary_key" json:"id"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
}

// ToPathWay ...
func ToPathWay(pathWayDTO *PathWayDTO) *PathWay {
	return &PathWay{
		Title: pathWayDTO.Title,
		Body:  pathWayDTO.Body,
	}
}

// ToPathWayDTO ...
func ToPathWayDTO(pathWay *PathWay) *PathWayDTO {
	return &PathWayDTO{
		ID:    pathWay.ID,
		Title: pathWay.Title,
		Body:  pathWay.Body,
	}
}
