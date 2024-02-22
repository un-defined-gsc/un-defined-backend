package social_repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
)

type TagsRepository struct {
	dbpool *pgxpool.Pool
}

func NewTagsRepository(dbpool *pgxpool.Pool) *TagsRepository {
	return &TagsRepository{
		dbpool: dbpool,
	}
}

func (r *TagsRepository) Create(ctx context.Context, tag *domains.CrateTagDTO) (err error) {
	query := `INSERT INTO t_tags (name,user_id,post_id) VALUES ($1,$2,$3)`
	_, err = r.dbpool.Exec(ctx, query, tag.Name, tag.UserID, tag.PostID)
	return
}
