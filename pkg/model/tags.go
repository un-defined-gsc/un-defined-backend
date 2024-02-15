package model

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id"`
	Tag       string    `json:"tag"`
}

type TagDTO struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `json:"username"`
	Tag      string    `json:"tag"`
}

func ToTag(tagDTO *TagDTO) *Tag {
	return &Tag{
		UserID: tagDTO.ID,
		Tag:    tagDTO.Tag,
	}
}

func ToTagDTO(tag *Tag) *TagDTO {
	return &TagDTO{
		ID:  tag.ID,
		Tag: tag.Tag,
	}
}
