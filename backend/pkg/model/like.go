package model

import (
	"time"

	"github.com/google/uuid"
)

// Like ...
type Like struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id"`
}

// LikeDTO ...
type LikeDTO struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `json:"username"`
}

// ToLike ...
func ToLike(likeDTO *LikeDTO) *Like {
	return &Like{
		UserID: likeDTO.ID,
	}
}

// ToLikeDTO ...
func ToLikeDTO(like *Like) *LikeDTO {
	return &LikeDTO{
		ID: like.ID,
	}
}
