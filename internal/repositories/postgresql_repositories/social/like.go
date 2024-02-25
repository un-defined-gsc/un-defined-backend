package social_repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type iLikesRepository struct {
	dbpool *pgxpool.Pool
}

func NewLikesRepository(dbpool *pgxpool.Pool) social_ports.ILikesRepository {
	return &iLikesRepository{
		dbpool: dbpool,
	}
}

func (r *iLikesRepository) Like(ctx context.Context, like *domains.LikeDTO) (err error) {
	_, err = r.dbpool.Exec(ctx, `
		INSERT INTO likes (user_id, post_id) VALUES ($1, $2)
	`, like.UserID, like.PostID)
	return
}

func (r *iLikesRepository) UnLike(ctx context.Context, like *domains.LikeDTO) (err error) {
	_, err = r.dbpool.Exec(ctx, `
		DELETE FROM likes WHERE post_id = $1, user_id = $2
	`, like.PostID, like.UserID)
	return
}

func (r *iLikesRepository) GetLikesByPostID(ctx context.Context, like *domains.LikeDTO) (likes []*domains.LikeDTO, err error) {
	rows, err := r.dbpool.Query(ctx, `
		SELECT user_id, post_id FROM likes WHERE post_id = $1 and user_id = $2
	`, like.PostID, like.UserID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		like := &domains.LikeDTO{}
		err = rows.Scan(&like.UserID, &like.PostID)
		if err != nil {
			return
		}
		likes = append(likes, like)
	}
	return

}
