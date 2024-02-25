package roadmap_ps_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type submissionRepository struct {
	dbpool *pgxpool.Pool
}

func NewSubmissionRepository(dbpool *pgxpool.Pool) roadmap_ports.ISubmissionRepository {
	return &submissionRepository{
		dbpool: dbpool,
	}
}

func (r *submissionRepository) Create(ctx context.Context, submission *roadmap_domain.Submission) (err error) {
	query := `INSERT INTO 
	t_submissions 
	(roadmap_id, user_id) VALUES ($1, $2) returning id`
	err = r.dbpool.QueryRow(ctx, query, submission.RoadmapID, submission.UserID).Scan(&submission.ID)
	return
}

func (r *submissionRepository) Delete(ctx context.Context, submissionID uuid.UUID) (err error) {
	query := `DELETE FROM t_submissions WHERE id = $1`
	_, err = r.dbpool.Exec(ctx, query, submissionID)
	return
}

func (r *submissionRepository) Filter(ctx context.Context, filter *roadmap_domain.Submission) (submissions []*roadmap_domain.Submission, err error) {
	query := `SELECT * FROM t_submissions WHERE
	($1::uuid  = uuid_nil() OR roadmap_id = $1) AND
	($3::uuid  = uuid_nil() OR user_id = $3)
	`
	rows, err := r.dbpool.Query(ctx, query, filter.RoadmapID, filter.UserID)
	if err != nil {
		return
	}
	defer rows.Close()
	submissions, err = pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[roadmap_domain.Submission])
	return
}
