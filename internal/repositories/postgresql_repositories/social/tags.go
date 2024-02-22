package social_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type tagsRepository struct {
	dbpool *pgxpool.Pool
}

func NewTagsRepository(dbpool *pgxpool.Pool) social_ports.ITagsRepository {
	return &tagsRepository{
		dbpool: dbpool,
	}
}

func (r *tagsRepository) Create(ctx context.Context, tag *domains.CrateTagDTO) (err error) {
	query := `INSERT INTO t_tags (name,user_id,post_id) VALUES ($1,$2,$3)`
	_, err = r.dbpool.Exec(ctx, query, tag.Name, tag.UserID, tag.PostID)
	return
}
func (r *tagsRepository) DeleteByID(ctx context.Context, tagID uuid.UUID) (err error) {
	query := `DELETE FROM t_tags WHERE id = $1`
	_, err = r.dbpool.Exec(ctx, query, tagID)
	return
}

func (r *tagsRepository) GetByID(ctx context.Context, postID uuid.UUID) ([]*domains.TagDTO, error) {
	query := `SELECT name FROM t_tags WHERE post_id = $1`
	rows, err := r.dbpool.Query(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*domains.TagDTO
	for rows.Next() {
		tag := new(domains.TagDTO)
		err := rows.Scan(&tag.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *tagsRepository) GetAll(ctx context.Context, limit, offset uint64) (tags []*domains.TagDTO, err error) {
	query := `SELECT name FROM t_tags ORDER BY created_at LIMIT $1 OFFSET $2 `
	rows, err := r.dbpool.Query(ctx, query, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		tag := new(domains.TagDTO)
		err = rows.Scan(&tag.Name)
		if err != nil {
			return
		}
		tags = append(tags, tag)
	}
	return
}
