package user_gorm_repositories

import (
	"context"

	"github.com/google/uuid"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
	gorm_migration "github.com/un-defined-gsc/un-defined-backend/internal/repositories/gorm_repositories"
	"gorm.io/gorm"
)

type bannedRepository struct {
	db *gorm.DB
}

func NewBannedsRepository(db *gorm.DB) user_ports.IBannedsRepository {
	repo := &bannedRepository{db: db}
	gorm_migration.Add(repo)
	return repo
}

func (r *bannedRepository) Create(ctx context.Context, banned *user_domain.Banned) (err error) {
	return
}

func (r *bannedRepository) Update(ctx context.Context, newBanned *user_domain.Banned) (err error) {
	return
}

func (r *bannedRepository) DeleteByID(ctx context.Context, bannedID uuid.UUID) (err error) {
	return
}

func (r *bannedRepository) DeleteByUserID(ctx context.Context, userID uuid.UUID) (err error) {
	return
}

func (r *bannedRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (banned *user_domain.Banned, err error) {
	return
}

func (r *bannedRepository) Migrate() (err error) {
	// Veri tabanı tablolarını oluştur
	err = r.db.AutoMigrate(&user_domain.Banned{})
	return
}
