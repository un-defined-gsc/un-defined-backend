package social_repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (r *iLikesRepository) Like(ctx context.Context, like *domains.LikeDTO) (likeID uuid.UUID, err error) {
	fmt.Println("Like")
	err = r.dbpool.QueryRow(ctx, `
		INSERT INTO t_likes (post_id, user_id) VALUES ($1, $2) returning id
	`, like.PostID, like.UserID).Scan(&likeID)
	return
}

func (r *iLikesRepository) UnLike(ctx context.Context, like *domains.LikeDTO) (likeID uuid.UUID, err error) {
	err = r.dbpool.QueryRow(ctx, `
		DELETE FROM t_likes WHERE post_id = $1 and user_id = $2 returning id
	`, like.PostID, like.UserID).Scan(&likeID)
	return
}

func (r *iLikesRepository) GetLikesByPostID(ctx context.Context, like *domains.LikeDTO) (likes []*domains.LikeDTO, err error) {
	rows, err := r.dbpool.Query(ctx, `
		SELECT user_id, post_id FROM t_likes WHERE post_id = $1 and user_id = $2
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
