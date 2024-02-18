package user_gorm_repositories

import (
	"context"

	"github.com/google/uuid"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
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
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return service_errors.ErrDataDuplication
		}
	}
	return
}

func (r *userRepository) Update(ctx context.Context, user *user_domain.User) (err error) {
	err = r.db.Exec(`UPDATE t_users SET
		first_name = $1,
		last_name = $2,
		gender = $4,
		appeal = $5,
		updated_at = CURRENT_TIMESTAMP,
		WHERE id = $4`, user.FirstName, user.LastName, user.Gender, user.Appeal, user.ID).Error

	return
}

func (r *userRepository) UpdateEmailVerifiedTrueByID(ctx context.Context, userID uuid.UUID) (err error) {
	err = r.db.Exec(`UPDATE t_users SET
			email_verified = true,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1`, userID).Error
	return
}

func (r *userRepository) UpdateDisabledByID(ctx context.Context, userID uuid.UUID, newDisabled bool) (err error) {
	err = r.db.Exec(`UPDATE t_users SET
			disabled = $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $2`, newDisabled, userID).Error

	return
}

func (r *userRepository) UpdateMasterAdminByID(ctx context.Context, userID uuid.UUID, newMasterAdmin bool) (err error) {
	err = r.db.Exec(`UPDATE t_users SET
		master_admin = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE uuid = $2`, newMasterAdmin, userID).Error
	return
}

func (r *userRepository) UpdatePasswordByID(ctx context.Context, userID uuid.UUID, newPassword string) (err error) {
	err = r.db.Exec(`
	UPDATE t_users SET
		password = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE uuid = $2`, newPassword, userID).Error
	return
}

func (r *userRepository) UpdateEmailByID(ctx context.Context, userID uuid.UUID, newEmail string) (err error) {
	err = r.db.Exec(`
	UPDATE t_users SET
		email = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE uuid = $2`, newEmail, userID).Error
	return
}

func (r *userRepository) UpdateLastLoginByID(ctx context.Context, userID uuid.UUID) (err error) {
	err = r.db.Exec(`
	UPDATE t_users SET
		last_login = CURRENT_TIMESTAMP,
	WHERE uuid = $1`, userID).Error
	return
}

func (r *userRepository) GetByID(ctx context.Context, userID uuid.UUID) (user *user_domain.User, err error) {
	err = r.db.Raw("SELECT * FROM t_users WHERE id = $1", userID).Scan(&user).Error
	return
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (user *user_domain.User, err error) {
	err = r.db.Raw("SELECT * FROM t_users WHERE email = $1", email).Scan(&user).Error
	return
}

func (r *userRepository) FindAll(ctx context.Context, searchUser user_domain.User, limit, offset uint64) (users []user_domain.User, count uint64, err error) {
	countQuery := `
	SELECT COUNT(*) FROM t_users
	WHERE
	($1::text = '' OR last_name ILIKE $1 || '%' )
	AND ($2::text = '' OR email ILIKE $2 || '%' )
	AND ($4::boolean IS NULL OR disabled = $4)
	AND ($5::boolean IS NULL OR master_admin = $5)
	AND ($6::boolean IS NULL OR banned = $6)
	AND ($7::boolean IS NULL OR email_verified = $7)
	AND ($8::boolean IS NULL OR mfa_enabled = $8)
	`
	err = r.db.Raw(countQuery, searchUser.FirstName, searchUser.LastName, searchUser.Email, searchUser.Disabled, searchUser.MasterAdmin, searchUser.Banned, searchUser.EmailVerified, searchUser.MFAEnabled).Scan(&count).Error
	if err != nil {
		return
	}

	query := `
	SELECT
		first_name, last_name, email, email_verified, banned,mfa_enabled, disabled,disabled_at,last_login,gender,appeal
	FROM t_users
	WHERE
	($1::text = '' OR last_name ILIKE $1 || '%' )
	AND ($2::text = '' OR email ILIKE $2 || '%' )
	AND ($4::boolean IS NULL OR disabled = $4)
	AND ($5::boolean IS NULL OR master_admin = $5)
	AND ($6::boolean IS NULL OR banned = $6)
	AND ($7::boolean IS NULL OR email_verified = $7)
	AND ($8::boolean IS NULL OR mfa_enabled = $8)
	ORDER BY created_at DESC
	LIMIT $9 OFFSET $10
	`
	err = r.db.Raw(query, searchUser.FirstName, searchUser.LastName, searchUser.Email, searchUser.Disabled, searchUser.MasterAdmin, searchUser.Banned, searchUser.EmailVerified, searchUser.MFAEnabled, limit, offset).Scan(&users).Error
	if err != nil {
		return
	}
	return
}

func (r *userRepository) GetCountByEmail(ctx context.Context, email string) (count uint64, err error) {
	err = r.db.Raw("SELECT COUNT(*) FROM t_users WHERE email = $1", email).Scan(&count).Error
	return
}

// IMigration interface
func (r *userRepository) Migrate() (err error) {
	// Veri tabanı tablolarını oluştur
	err = r.db.AutoMigrate(&user_domain.User{})
	return
}
