package user_gorm_repositories

import (
	"context"

	"github.com/google/uuid"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
	gorm_migration "github.com/un-defined-gsc/un-defined-backend/internal/repositories/gorm_repositories"
	"gorm.io/gorm"
)

type mfaRepository struct {
	db *gorm.DB
}

func NewMFAsRepository(db *gorm.DB) user_ports.IMFAsRepository {
	repo := &mfaRepository{db: db}
	gorm_migration.Add(repo)
	return repo
}

func (r *mfaRepository) Create(ctx context.Context, mfaSetting *user_domain.MFASetting) (err error) {
	return
}

func (r *mfaRepository) Update(ctx context.Context, mfaSetting *user_domain.MFASetting) (err error) {
	return
}

func (r *mfaRepository) DeleteByUserID(ctx context.Context, userID uuid.UUID) (err error) {
	return
}

func (r *mfaRepository) UpdateLogIDByID(ctx context.Context, mfaSettingID, newLogID uuid.UUID) (err error) {
	return
}

func (r *mfaRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (mfaSetting *user_domain.MFASetting, err error) {
	return
}

func (r *mfaRepository) Migrate() (err error) {
	// Veri tabanı tablolarını oluştur
	err = r.db.AutoMigrate(&user_domain.MFASetting{})
	return
}
