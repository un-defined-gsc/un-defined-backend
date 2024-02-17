package user_gorm_repositories

import (
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

func (r *mfaRepository) Migrate() error {
	// Veri tabanı tablolarını oluştur
	return r.db.AutoMigrate(&user_domain.MFASetting{})
}
