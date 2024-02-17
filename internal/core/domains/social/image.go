package social_domain

import "github.com/google/uuid"

type Image struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	PostID    uuid.UUID `json:"post_id"`
	UserID    uuid.UUID `json:"user_id"`
	Path      string    `json:"path"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type ImageDTO struct {
	ID     uuid.UUID `gorm:"primary_key" json:"id"`
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
	Path   string    `json:"path"`
}

func ToImage(imageDTO *ImageDTO) *Image {
	return &Image{
		PostID: imageDTO.PostID,
		Path:   imageDTO.Path,
	}
}

func ToImageDTO(image *Image) *ImageDTO {
	return &ImageDTO{
		ID:     image.ID,
		PostID: image.PostID,
		UserID: image.UserID,
		Path:   image.Path,
	}
}
