package user_ps_repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

type usersRepository struct {
	dbpool *pgxpool.Pool
}

func NewUsersRepository(dbpool *pgxpool.Pool) user_ports.IUsersRepository {
	return &usersRepository{
		dbpool: dbpool,
	}
}

func (r *usersRepository) Create(ctx context.Context, user *user_domain.User) (err error) {
	query := `
		INSERT INTO t_users (
			password,
			first_name,
			last_name,
			lang,
			email,
			email_verified,
			gender,
			appeal
			) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8) returning id`
	err = r.dbpool.QueryRow(ctx, query, user.Password, user.FirstName, user.LastName, user.Lang, user.Email, user.EmailVerified, user.Gender, user.Appeal).Scan(&user.ID)
	if err != nil {
		pgErr := &pgconn.PgError{}
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return service_errors.ErrDataDuplication
			}
		}
	}
	return
}

func (r *usersRepository) Update(ctx context.Context, newUser *user_domain.User) (err error) {
	query := `
		UPDATE t_users SET
			first_name = $1,
			last_name = $2,
			lang = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $4`
	_, err = r.dbpool.Exec(ctx, query, newUser.FirstName, newUser.LastName, newUser.Lang, newUser.ID)
	return
}

func (r *usersRepository) UpdateEmailVerifiedTrueByID(ctx context.Context, userID uuid.UUID) (err error) {
	query := `
		UPDATE t_users SET
			email_verified = true,
			updated_at = now()
		WHERE id = $1`
	_, err = r.dbpool.Exec(ctx, query, userID)

	return
}

func (r *usersRepository) UpdateDisabledByID(ctx context.Context, userID uuid.UUID, newDisabled bool) (err error) {
	query := `
		UPDATE t_users SET
			disabled = $1,
			updated_at = now()
		WHERE id = $2`
	_, err = r.dbpool.Exec(ctx, query, newDisabled, userID)

	return
}

func (r *usersRepository) UpdateMasterAdminByID(ctx context.Context, userID uuid.UUID, newMasterAdmin bool) (err error) {
	query := `
		UPDATE t_users SET
			master_admin = $1,
			updated_at = now()
		WHERE id = $2`
	_, err = r.dbpool.Exec(ctx, query, newMasterAdmin, userID)

	return
}

func (r *usersRepository) UpdatePasswordByID(ctx context.Context, userID uuid.UUID, newPassword string) (err error) {
	query := `
		UPDATE t_users SET
			password = $1,
			updated_at = now()
		WHERE id = $2`
	_, err = r.dbpool.Exec(ctx, query, newPassword, userID)

	return
}

func (r *usersRepository) UpdateLastLoginByID(ctx context.Context, userID uuid.UUID) (err error) {
	query := `
		UPDATE t_users SET
			last_login = now()
			WHERE id = $1`
	_, err = r.dbpool.Exec(ctx, query, userID)
	return
}

func (r *usersRepository) GetByID(ctx context.Context, userID uuid.UUID) (user *user_domain.User, err error) {
	query := `
		SELECT
			*
		FROM t_users
		WHERE id = $1`
	rows, err := r.dbpool.Query(ctx, query, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, service_errors.ErrDataNotFound
		}
		return
	}
	defer rows.Close()
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[user_domain.User])
}

func (r *usersRepository) GetByEmail(ctx context.Context, email string) (user *user_domain.User, err error) {
	query := `
		SELECT
			*
		FROM t_users
		WHERE email = $1`
	row, err := r.dbpool.Query(ctx, query, email)
	if err != nil {
		return
	}
	defer row.Close()
	user, err = pgx.CollectOneRow(row, pgx.RowToAddrOfStructByName[user_domain.User])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, service_errors.ErrDataNotFound
		}
	}
	return
}

func (r *usersRepository) FindAll(ctx context.Context, searchUser user_domain.User, limit, offset uint64) (users []user_domain.User, count uint64, err error) {
	query := `
		SELECT
			*, COUNT(*) OVER() AS total_count
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
		`
	rows, err := r.dbpool.Query(ctx, query, searchUser.FirstName, searchUser.LastName, searchUser.Email, searchUser.Disabled, searchUser.MasterAdmin, searchUser.Banned, searchUser.EmailVerified, searchUser.MFAEnabled, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&count)
	if err != nil {
		return
	}
	users, err = pgx.CollectRows(rows, pgx.RowToStructByName[user_domain.User])
	return
}

func (r *usersRepository) UpdateEmailByID(ctx context.Context, userID uuid.UUID, newEmail string) (err error) {
	query := `
		UPDATE t_users SET
			email = $1,
			email_verified = true,
			updated_at = now()
		WHERE id = $2`
	_, err = r.dbpool.Exec(ctx, query, newEmail, userID)
	return
}

func (r *usersRepository) GetCountByEmail(ctx context.Context, email string) (count uint64, err error) {
	query := `
		SELECT COUNT(*)
		FROM t_users
		WHERE email = $1`
	row, err := r.dbpool.Query(ctx, query, email)
	if err != nil {
		return
	}
	defer row.Close()
	row.Next()
	err = row.Scan(&count)
	if err != nil {
		return
	}
	return
}
