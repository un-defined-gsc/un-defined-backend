package user_gorm_repositories

import (
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
	"gorm.io/gorm"
)

type mfaRepository struct {
	db *gorm.DB
}

func NewMFAsRepository(db *gorm.DB) user_ports.IMFAsRepository {
	return &mfaRepository{db: db}
}
