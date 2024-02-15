package model

import (
	"time"

	"github.com/google/uuid"
)

// Comment ...
type Comment struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id"`
	Body      string    `json:"body"`
}

// CommentDTO ...
type CommentDTO struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `json:"username"`
	PostID   uuid.UUID `json:"post_id"`
	Body     string    `json:"body"`
}

// ToComment ...
func ToComment(commentDTO *CommentDTO) *Comment {
	return &Comment{
		UserID: commentDTO.ID,
		Body:   commentDTO.Body,
	}
}

// ToCommentDTO ...
func ToCommentDTO(comment *Comment) *CommentDTO {
	return &CommentDTO{
		ID:     comment.ID,
		Body:   comment.Body,
		PostID: comment.PostID,
	}
}
