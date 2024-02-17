package user_gorm_repositories

import (
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
	"gorm.io/gorm"
)

type bannedRepository struct {
	db *gorm.DB
}

func NewBannedsRepository(db *gorm.DB) user_ports.IBannedsRepository {
	return &bannedRepository{db: db}
}
