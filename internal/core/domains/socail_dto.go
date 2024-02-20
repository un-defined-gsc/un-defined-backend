package domains

import (
	"time"

	"github.com/google/uuid"
)

//----- Social ------//

type CategoryDTO struct {
	ID   uuid.UUID `gorm:"primary_key" json:"id"`
	Name string    `json:"category"`
}

type PostDTO struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Likes     uint64    `json:"likes"`
	Comments  uint64    `json:"comments"`
	Images    []string  `json:"images"`
	Tags      []TagDTO  `json:"tags"`
}

type InPostDTO struct {
	Category  string       `json:"category"`
	Name      string       `json:"name"`
	Surname   string       `json:"surname"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	Likes     uint64       `json:"likes"`
	Comments  []CommentDTO `json:"comments"`
	Images    []string     `json:"images"`
	Tags      []TagDTO     `json:"tags"`
	CreatedAt time.Time    `json:"created_at"`
}

type CommentDTO struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
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

type CratePostDTO struct {
	ID       uuid.UUID `json:"-"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Category string    `json:"category"`
	UserID   uuid.UUID `json:"-"`
	Tags     []string  `json:"tags"`
	Image    []string  `json:"image"`
}

type UpdatePostDTO struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Category string    `json:"category"`
	Tags     []string  `json:"tags"`
	Image    []string  `json:"image"`
}
