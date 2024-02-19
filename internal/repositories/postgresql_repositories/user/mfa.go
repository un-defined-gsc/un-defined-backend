package user_ps_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

// Yeniden yazılacak
// KeyType panic yaptırabilir
type mfaSettingRepository struct {
	db *pgxpool.Pool
}

func NewMFAsRepository(db *pgxpool.Pool) user_ports.IMFAsRepository {
	return &mfaSettingRepository{db: db}
}

// last_log_uuid şimdilik boş bırakıldı çünkü loglama servisi yazılmadı
func (p *mfaSettingRepository) Create(ctx context.Context, mfaSetting *user_domain.MFASetting) (err error) {
	tx, err := p.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	query := "UPDATE t_users SET mfa_enabled = TRUE WHERE id = $1"
	rows, err := tx.Exec(ctx, query, mfaSetting.UserID)
	if err != nil {
		return
	}
	if rows.RowsAffected() <= 0 {
		return pgx.ErrNoRows
	}
	query = "INSERT INTO t_mfa_settings (user_id,key,created_at) VALUES ($1,$2,$3)" // last_log_uuid = $5
	_, err = p.db.Exec(ctx, query, mfaSetting.UserID, mfaSetting.Key, mfaSetting.CreatedAt)
	if err != nil {
		return
	}
	err = tx.Commit(ctx)
	return
}

func (p *mfaSettingRepository) Update(ctx context.Context, newMFASetting *user_domain.MFASetting) (err error) {
	query := "UPDATE t_mfa_settings SET key = $1, created_at = $2, WHERE user_id = $3" // last_log_uuid = $4 , user_uuid = $5
	_, err = p.db.Exec(ctx, query, newMFASetting.Key, newMFASetting.CreatedAt, newMFASetting.UserID)
	return
}

func (p *mfaSettingRepository) DeleteByUserID(ctx context.Context, userID uuid.UUID) (err error) {
	tx, err := p.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}
	defer tx.Rollback(ctx)
	query := "UPDATE t_users SET mfa_enabled = FALSE WHERE id = $1"
	rows, err := tx.Exec(ctx, query, userID)
	if err != nil {
		return
	}
	if rows.RowsAffected() <= 0 {
		return pgx.ErrNoRows
	}
	query = "DELETE FROM t_mfa_settings WHERE user_id = $1"
	_, err = p.db.Exec(ctx, query, userID)
	if err != nil {
		return
	}
	err = tx.Commit(ctx)
	return
}
func (p *mfaSettingRepository) UpdateLogIDByID(ctx context.Context, mfaSettingID, newLogID uuid.UUID) (err error) { // Bu fonksiyonu kullanmıyoruz
	query := "UPDATE t_mfa_settings SET last_log_uuid = $1 WHERE user_id = $2"
	_, err = p.db.Exec(ctx, query, newLogID, mfaSettingID)
	return
}

func (p *mfaSettingRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (mfaSetting *user_domain.MFASetting, err error) {
	query := "SELECT * FROM t_mfa_settings WHERE user_id = $1"
	rows, err := p.db.Query(ctx, query, userID)
	if err != nil {
		return
	}
	mfaSetting, err = pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[user_domain.MFASetting])
	if err != nil {
		if err == pgx.ErrNoRows {
			err = service_errors.ErrDataNotFound
		}
		return
	}
	return
}
