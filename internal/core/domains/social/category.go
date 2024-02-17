package social_domain

import (
	"time"

	"github.com/google/uuid"
)

// Category ...
type Category struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"category"`
}

// CategoryDTO ...
type CategoryDTO struct {
	ID   uuid.UUID `gorm:"primary_key" json:"id"`
	Name string    `json:"category"`
}

// ToCategory ...
func ToCategory(categoryDTO *CategoryDTO) *Category {
	return &Category{
		Name: categoryDTO.Name,
	}
}

// ToCategoryDTO ...
func ToCategoryDTO(category *Category) *CategoryDTO {
	return &CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}
