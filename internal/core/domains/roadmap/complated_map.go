package roadmap_domain

import (
	"time"

	"github.com/google/uuid"
)

// CompletedMap ...
type CompletedMap struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	PathWayID uuid.UUID `json:"pathway_id"`
	RoadmapID uuid.UUID `json:"roadmap_id"`
	SubPathID uuid.UUID `json:"subpath_id"`
}

// CompletedMapDTO ...
type CompletedMapDTO struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	PathWayID uuid.UUID `json:"pathway_id"`
	RoadmapID uuid.UUID `json:"roadmap_id"`
	SubPathID uuid.UUID `json:"subpath_id"`
}

// ToCompletedMap ...
func ToCompletedMap(completedMapDTO *CompletedMapDTO) *CompletedMap {
	return &CompletedMap{
		UserID:    completedMapDTO.UserID,
		PathWayID: completedMapDTO.PathWayID,
		RoadmapID: completedMapDTO.RoadmapID,
		SubPathID: completedMapDTO.SubPathID,
	}
}

// ToCompletedMapDTO ...
func ToCompletedMapDTO(completedMap *CompletedMap) *CompletedMapDTO {
	return &CompletedMapDTO{
		ID:        completedMap.ID,
		UserID:    completedMap.UserID,
		PathWayID: completedMap.PathWayID,
		RoadmapID: completedMap.RoadmapID,
		SubPathID: completedMap.SubPathID,
	}
}
