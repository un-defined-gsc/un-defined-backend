package service

import (
	"un-defined/pkg/model"

	repuser "un-defined/pkg/repository/user"

	"github.com/google/uuid"
)

// UserService ...
type UserService struct {
	UserRepository *repuser.Repository
}

// NewUserService ...
func NewUserService(u *repuser.Repository) UserService {
	return UserService{UserRepository: u}
}

// All ...
func (u *UserService) All() ([]model.User, error) {
	return u.UserRepository.All()
}

// FindByID ...
func (u *UserService) FindByID(id uint) (*model.User, error) {
	return u.UserRepository.FindByID(id)
}

// Save ...
func (u *UserService) Save(user *model.User) (*model.User, error) {
	return u.UserRepository.Save(user)
}

// Delete ...
func (u *UserService) Delete(id uuid.UUID) error {
	return u.UserRepository.Delete(id)
}

// Migrate ...
func (u *UserService) Migrate() error {
	return u.UserRepository.Migrate()
}
