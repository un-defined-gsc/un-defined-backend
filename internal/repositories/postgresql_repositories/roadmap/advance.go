package roadmap_ps_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type advanceRepository struct {
	dbpool *pgxpool.Pool
}

func NewAdvanceRepository(dbpool *pgxpool.Pool) roadmap_ports.IAdvanceRepository {
	return &advanceRepository{
		dbpool: dbpool,
	}
}

func (r *advanceRepository) Create(ctx context.Context, advance *roadmap_domain.Advance) (err error) {
	query := `INSERT INTO t_advance (roadmap_id,user_id,pathway_id,advance_type) VALUES ($1,$2,$3,$4) returning id`
	err = r.dbpool.QueryRow(ctx, query, advance.RoadmapID, advance.UserID, advance.PathWayID, advance.AdvanceType).Scan(&advance.ID)
	return
}

func (r *advanceRepository) Delete(ctx context.Context, advanceId uuid.UUID) (err error) {
	query := `DELETE FROM t_advance WHERE id = $1`
	_, err = r.dbpool.Exec(ctx, query, advanceId)
	return
}

func (r *advanceRepository) Filter(ctx context.Context, filter *roadmap_domain.Advance) (advances []*roadmap_domain.Advance, err error) {
	query := `SELECT * FROM t_advance WHERE
		($1::uuid IS uuid_nil() OR roadmap_id = $1) AND
		($2::uuid IS uuid_nil() OR user_id = $2) AND
		($3::uuid IS uuid_nil() OR pathway_id = $3) AND
		($4::text IS "" OR advance_type = ILIKE $4 || '%')
		`
	rows, err := r.dbpool.Query(ctx, query, filter.RoadmapID, filter.UserID, filter.PathWayID, filter.AdvanceType)
	if err != nil {
		return
	}
	advances, err = pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[roadmap_domain.Advance])
	return
}
