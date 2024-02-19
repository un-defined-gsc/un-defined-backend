package user_ps_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

type bannedsRepository struct {
	db *pgxpool.Pool
}

func NewBannedsRepository(db *pgxpool.Pool) user_ports.IBannedsRepository {
	return &bannedsRepository{db: db}
}

// log_uuid şimdilik boş bırakıldı çünkü loglama servisi yazılmadı
func (p *bannedsRepository) Create(ctx context.Context, banned *user_domain.Banned) (err error) {
	tx, err := p.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}
	defer func() {
		err = tx.Rollback(ctx)
	}()
	query := "UPDATE t_users SET banned = TRUE WHERE id = $1"
	rows, err := tx.Exec(ctx, query, banned.ID)
	if err != nil {
		return
	}
	if rows.RowsAffected() <= 0 {
		return pgx.ErrNoRows
	}
	query = "INSERT INTO t_banned (user_id,admin_id,reason,expires_at,permanent) VALUES ($1,$2,$3,$4,$5)"
	_, err = tx.Exec(ctx, query, banned.ID, banned.ID, banned.Reason, banned.ExpiresAt, banned.Permanent)
	if err != nil {
		return
	}
	err = tx.Commit(ctx)
	return
}

func (p *bannedsRepository) Update(ctx context.Context, newBanned *user_domain.Banned) (err error) {
	query := "UPDATE t_banned SET admin_id = $1, reason = $2, expires_at = $3, permanent = $4 WHERE id = $5"
	_, err = p.db.Exec(ctx, query, newBanned.AdminID, newBanned.Reason, newBanned.ExpiresAt, newBanned.Permanent, newBanned.AdminID)
	return
}

func (p *bannedsRepository) UpdateLogIDByID(ctx context.Context, bannedUUID, newLogUUID uuid.UUID) (err error) {
	query := "UPDATE t_banned SET log_uuid = $1 WHERE id = $2"
	_, err = p.db.Exec(ctx, query, newLogUUID, bannedUUID)
	return
}

func (p *bannedsRepository) DeleteByID(ctx context.Context, bannedUUID uuid.UUID) (err error) {
	query := "DELETE FROM t_banned WHERE id = $1"
	_, err = p.db.Exec(ctx, query, bannedUUID)
	return
}

func (p *bannedsRepository) DeleteByUserID(ctx context.Context, userUUID uuid.UUID) (err error) {
	tx, err := p.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}
	defer func() {
		err = tx.Rollback(ctx)
	}()
	query := "UPDATE t_users SET banned = FALSE WHERE id = $1"
	rows, err := tx.Exec(ctx, query, userUUID)
	if err != nil {
		return
	}
	if rows.RowsAffected() <= 0 {
		return pgx.ErrNoRows
	}
	query = "DELETE FROM t_banned WHERE user_id = $1"
	_, err = tx.Exec(ctx, query, userUUID)
	if err != nil {
		return
	}
	err = tx.Commit(ctx)
	return
}

func (p *bannedsRepository) GetByUserID(ctx context.Context, userUUID uuid.UUID) (banned *user_domain.Banned, err error) {
	query := "SELECT * FROM t_banned WHERE user_id = $1"
	rows, err := p.db.Query(ctx, query, userUUID)
	if err != nil {
		return
	}
	defer rows.Close()
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[user_domain.Banned])

}
