package user_gorm_repositories

import (
	"context"

	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
	gorm_migration "github.com/un-defined-gsc/un-defined-backend/internal/repositories/gorm_repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) user_ports.IUsersRepository {
	repo := &userRepository{db: db}
	gorm_migration.Add(repo)
	return repo
}

// IUser
func (r *userRepository) Create(ctx context.Context, user *user_domain.User) (err error) {

}

// IMigration interface
func (r *userRepository) Migrate() error {
	// Veri tabanı tablolarını oluştur
	return r.db.AutoMigrate(&user_domain.User{})
}
