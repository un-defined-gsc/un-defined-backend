package domains

import (
	"time"

	"github.com/google/uuid"
)

//----- Social ------//

type CategoryDTO struct {
	Name string `json:"category"`
}

type PostDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Likes     uint64    `json:"likes"`
	Comments  uint64    `json:"comments"`
	Images    []string  `json:"images"`
	Tags      []TagDTO  `json:"tags"`
}

type InPostDTO struct {
	Category   string       `json:"category"`
	Name       string       `json:"name"`
	Surname    string       `json:"surname"`
	Title      string       `json:"title"`
	Content    string       `json:"content"`
	Likes      uint64       `json:"likes"`
	Editable   bool         `json:"editable"`
	Deleteable bool         `json:"deleteable"`
	Comments   []CommentDTO `json:"comments"`
	Images     []string     `json:"images"`
	Tags       []TagDTO     `json:"tags"`
	CreatedAt  time.Time    `json:"created_at"`
}
type CratePostDTO struct {
	ID       uuid.UUID `json:"-" `
	Title    string    `json:"title" validate:"required" example:"title"`
	Content  string    `json:"content" validate:"required" example:"content"`
	Category string    `json:"category" validate:"required,oneof=question problem story jobadvert" example:"story"`
	UserID   uuid.UUID `json:"-"`
	Tags     []string  `json:"tags" example:"tags"`
	Image    []string  `json:"image" example:"image"`
}
type CommentDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type TagDTO struct {
	Name string `json:"tag"`
}

type PostsDTO struct {
	Posts []PostDTO `json:"posts"`
}

type UpdatePostDTO struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Title    string    `json:"title"`
	UserID   uuid.UUID `json:"-"`
	Content  string    `json:"content"`
	Category string    `json:"category"`
	Tags     []string  `json:"tags"`
	Image    []string  `json:"image"`
}

type CrateTagDTO struct {
	Name   string    `json:"name" validate:"required"`
	UserID uuid.UUID `json:"-"`
	PostID uuid.UUID `json:"-"`
}
