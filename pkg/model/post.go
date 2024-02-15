package model

import (
	"time"

	"github.com/google/uuid"
)

// Post is our main model for Posts
type Post struct {
	ID         uuid.UUID `gorm:"primary_key" json:"id"`
	CategoryID uuid.UUID `json:"category_id"`
	UserID     uuid.UUID `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	
}

// PostDTO is our data transfer object for Post
type PostDTO struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Username string    `json:"username"`
	Category string    `json:"category"`
}

// ToPost converts postDTO to post
func ToPost(postDTO *PostDTO) *Post {
	return &Post{
		Title: postDTO.Title,
		Body:  postDTO.Body,
	}
}

// ToPostDTO converts post to postDTO
func ToPostDTO(post *Post) *PostDTO {
	return &PostDTO{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}
}

/* Example JSON
{
	"Title":"Dummy Title",
	"Body":"Dummy content",
}
*/
