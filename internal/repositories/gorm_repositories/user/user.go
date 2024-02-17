package user_gorm_repositories

import (
	"context"

	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user_ports.IUsersRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *user_domain.User) (err error) {
	r.db.WithContext(ctx)