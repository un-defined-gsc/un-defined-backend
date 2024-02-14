package model

import (
	"time"

	"github.com/google/uuid"
)

// Roadmap ...
type Roadmap struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}

// RoadmapDTO ...
type RoadmapDTO struct {
	ID    uuid.UUID `gorm:"primary_key" json:"id"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
}

// ToRoadmap ...
func ToRoadmap(roadmapDTO *RoadmapDTO) *Roadmap {
	return &Roadmap{
		Title: roadmapDTO.Title,
		Body:  roadmapDTO.Body,
	}
}

// ToRoadmapDTO ...
func ToRoadmapDTO(roadmap *Roadmap) *RoadmapDTO {
	return &RoadmapDTO{
		ID:    roadmap.ID,
		Title: roadmap.Title,
		Body:  roadmap.Body,
	}
}
