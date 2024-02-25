package social_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type iCommentsRepository struct {
	dbpool *pgxpool.Pool
}

func NewCommentsRepository(dbpool *pgxpool.Pool) social_ports.ICommentsRepository {
	return &iCommentsRepository{
		dbpool: dbpool,
	}
}

func (r *iCommentsRepository) Create(ctx context.Context, comment *domains.CommentDTO) (err error) {
	_, err = r.dbpool.Exec(ctx, `
		INSERT INTO t_comments (user_id, post_id, comments) VALUES ($1, $2, $3)
	`, comment.UserID, comment.PostID, comment.Body)
	return
}

func (r *iCommentsRepository) DeleteByID(ctx context.Context, commentID, userID uuid.UUID) (err error) {
	_, err = r.dbpool.Exec(ctx, `
		DELETE FROM t_comments WHERE id = $1 and user_id = $2
	`, commentID, userID)
	return
}

func (r *iCommentsRepository) GetAllByPostID(ctx context.Context, postID uuid.UUID, limit, offset uint64) (comments []domains.ResCommentDTO, err error) {
	rows, err := r.dbpool.Query(ctx, `
		SELECT c.id, u.first_name, u.last_name, c.comments, c.created_at FROM t_comments c
		INNER JOIN t_users u ON c.user_id = u.id
		WHERE c.post_id = $1
		LIMIT $2 OFFSET $3
	`, postID, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		comment := &domains.ResCommentDTO{}
		err = rows.Scan(&comment.ID, &comment.Name, &comment.Surname, &comment.Body, &comment.CreatedAt)
		if err != nil {
			return
		}
		comments = append(comments, *comment)
	}
	return
}
