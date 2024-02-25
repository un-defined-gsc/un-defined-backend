package roadmap_ps_repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type roadmapRepository struct {
	dbpool *pgxpool.Pool
}

func NewRoadmapRepository(dbpool *pgxpool.Pool) roadmap_ports.IRoadmapRepository {
	return &roadmapRepository{
		dbpool: dbpool,
	}
}

func (r *roadmapRepository) Create(ctx context.Context, roadmap *roadmap_domain.Roadmap) (err error) {
	query := `INSERT INTO 
	t_roadmaps 
	(name, description, first_path_id) VALUES ($1, $2, $3) returning id`
	err = r.dbpool.QueryRow(ctx, query, roadmap.Name, roadmap.Description, roadmap.FirstPathID).Scan(&roadmap.ID)
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

func (r *roadmapRepository) Update(ctx context.Context, newRoadmap *roadmap_domain.Roadmap) (err error) {
	query := `UPDATE t_roadmaps SET name = $1, description = $2, first_path_id = $3 WHERE id = $4`
	_, err = r.dbpool.Exec(ctx, query, newRoadmap.Name, newRoadmap.Description, newRoadmap.FirstPathID, newRoadmap.ID)
	return
}

func (r *roadmapRepository) Delete(ctx context.Context, roadmapID uuid.UUID) (err error) {
	query := `DELETE FROM t_roadmaps WHERE id = $1`
	_, err = r.dbpool.Exec(ctx, query, roadmapID)
	return
}

func (r *roadmapRepository) Filter(ctx context.Context, filter *roadmap_domain.Roadmap) (roadmaps []*roadmap_domain.Roadmap, err error) {
	query := `SELECT * FROM t_roadmaps WHERE
	($1::uuid IS uuid_nil() OR id = $1) AND
	($2::text IS NULL OR name ILIKE $2 || '%') AND
	($3::text IS NULL OR description ILIKE $3 || '%') AND
	($4::uuid IS NULL OR first_path_id = $4)
	`
	rows, err := r.dbpool.Query(ctx, query, filter.ID, filter.Name, filter.Description, filter.FirstPathID)
	if err != nil {
		return
	}
	roadmaps, err = pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[roadmap_domain.Roadmap])
	return nil, nil
}
