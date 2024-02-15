package service

import (
	"un-defined/pkg/model"
	"un-defined/pkg/repository/post"

	"github.com/google/uuid"
)

// PostService ...
type PostService struct {
	PostRepository *post.Repository
}

// NewPostService ...
func NewPostService(p *post.Repository) PostService {
	return PostService{PostRepository: p}
}

// All ...
func (p *PostService) All() ([]model.Post, error) {
	return p.PostRepository.All()
}

// FindByID ...
func (p *PostService) FindByID(id uint) (*model.Post, error) {
	return p.PostRepository.FindByID(id)
}

// Save ...
func (p *PostService) Save(post *model.Post) (*model.Post, error) {
	
	return p.PostRepository.Save(post)
}

// Delete ...
func (p *PostService) Delete(id uuid.UUID) error {
	return p.PostRepository.Delete(id)
}

// Migrate ...
func (p *PostService) Migrate() error {
	return p.PostRepository.Migrate()
}
