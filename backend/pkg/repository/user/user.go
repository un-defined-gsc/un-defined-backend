package repository

import (
	"un-defined/pkg/model"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// UserRepository ...
type Repository struct {
	db *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// All ...
func (u *Repository) All() ([]model.User, error) {
	users := []model.User{}
	err := u.db.Find(&users).Error
	return users, err
}

// FindByID ...
func (u *Repository) FindByID(id uint) (*model.User, error) {
	user := new(model.User)
	err := u.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

// Save ...
func (u *Repository) Save(user *model.User) (*model.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

// Delete ...
func (u *Repository) Delete(id uuid.UUID) error {
	err := u.db.Delete(&model.User{ID: id}).Error
	return err
}

// Migrate ...
func (u *Repository) Migrate() error {
	return u.db.AutoMigrate(&model.User{}).Error
}
