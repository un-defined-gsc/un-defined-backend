package user_gorm_repositories

import (
	"context"

	"github.com/google/uuid"
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
	err = r.db.Raw("INSERT INTO t_users (password, first_name, last_name, lang, email, email_verified, gender, appeal ) VALUES ( $1, $2, $3, $4, $5, $6 ,$7, $8) returning id",
		user.Password,
		user.FirstName,
		user.LastName,
		user.Lang,
		user.Email,
		user.EmailVerified,
		user.Gender,
		user.Appeal,
	).Scan(&user.ID).Error
	return
}

func (r *userRepository) Update(ctx context.Context, user *user_domain.User) (err error) {
	return
}

func (r *userRepository) UpdateEmailVerifiedTrueByID(ctx context.Context, userID uuid.UUID) (err error) {
	return
}

func (r *userRepository) UpdateDisabledByID(ctx context.Context, userID uuid.UUID, newDisabled bool) (err error) {
	return
}

func (r *userRepository) UpdateMasterAdminByID(ctx context.Context, userID uuid.UUID, newMasterAdmin bool) (err error) {
	return
}

func (r *userRepository) UpdatePasswordByID(ctx context.Context, userID uuid.UUID, newPassword string) (err error) {
	return
}

func (r *userRepository) UpdateEmailByID(ctx context.Context, userID uuid.UUID, newEmail string) (err error) {
	return
}

func (r *userRepository) UpdateLastLoginByID(ctx context.Context, userID uuid.UUID) (err error) {
	return
}

func (r *userRepository) GetByID(ctx context.Context, userID uuid.UUID) (user *user_domain.User, err error) {
	return
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (user *user_domain.User, err error) {
	return
}

func (r *userRepository) FindAll(ctx context.Context, searchUser user_domain.User, limit, offset uint64) (users []user_domain.User, count uint64, err error) {
	return
}

func (r *userRepository) GetCountByEmail(ctx context.Context, email string) (count uint64, err error) {
	return
}

// IMigration interface
func (r *userRepository) Migrate() (err error) {
	// Veri tabanı tablolarını oluştur
	err = r.db.AutoMigrate(&user_domain.User{})
	return
}
