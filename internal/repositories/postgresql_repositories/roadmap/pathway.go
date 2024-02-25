package roadmap_ps_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type pathWayRepository struct {
	dbpool *pgxpool.Pool
}

func NewPathWayRepository(dbpool *pgxpool.Pool) roadmap_ports.IPathWayRepository {
	return &pathWayRepository{
		dbpool: dbpool,
	}
}

func (r *pathWayRepository) Create(ctx context.Context, pathway *roadmap_domain.PathWay) (err error) {
	query := `INSERT INTO 
	t_pathways 
	(roadmap_id, name, description, parent_id) VALUES ($1, $2, $3, $4) returning id`
	err = r.dbpool.QueryRow(ctx, query, pathway.RoadmapID, pathway.Name, pathway.Description, pathway.ParentID).Scan(&pathway.ID)
	return
}

func (r *pathWayRepository) Update(ctx context.Context, newPathway *roadmap_domain.PathWay) (err error) {
	query := `UPDATE t_pathways SET name = $1, description = $2, parent_id = $3 WHERE id = $4`
	_, err = r.dbpool.Exec(ctx, query, newPathway.Name, newPathway.Description, newPathway.ParentID, newPathway.ID)
	return
}

func (r *pathWayRepository) Delete(ctx context.Context, pathWayID uuid.UUID) (err error) {
	query := `DELETE FROM t_pathways WHERE id = $1`
	_, err = r.dbpool.Exec(ctx, query, pathWayID)
	return
}

func (r *pathWayRepository) Filter(ctx context.Context, filter *roadmap_domain.PathWay) (pathways []*roadmap_domain.PathWay, err error) {
	query := `SELECT * FROM t_pathways WHERE
	($1::uuid IS uuid_nil() OR roadmap_id = $1) AND
	($2::text IS NULL OR name ILIKE $2 || '%') AND
	($3::text IS NULL OR description ILIKE $2 || '%') AND
	($4::uuid IS uuid_nil() OR parent_id = $4)
	`
	rows, err := r.dbpool.Query(ctx, query, filter.RoadmapID, filter.Name, filter.Description, filter.ParentID)
	if err != nil {
		return
	}
	pathways, err = pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[roadmap_domain.PathWay])
	return
}
