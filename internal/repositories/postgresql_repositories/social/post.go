package social_repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type postRepository struct {
	dbpool *pgxpool.Pool
}

func NewPostRepository(dbpool *pgxpool.Pool) social_ports.IPostsRepository {
	return &postRepository{
		dbpool: dbpool,
	}
}

func (r *postRepository) Create(ctx context.Context, post *domains.CratePostDTO, categoryID uuid.UUID) (postID uuid.UUID, err error) {
	query := `
		INSERT INTO t_posts (
			title,
			content,
			category_id,
			user_id
			) VALUES ( $1, $2, $3, $4) returning id`
	err = r.dbpool.QueryRow(ctx, query, post.Title, post.Content, categoryID, post.UserID).Scan(&post.ID)
	if err != nil {
		pgErr := &pgconn.PgError{}
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return uuid.Nil, service_errors.ErrDataDuplication
			}
		}
	}
	return
}
func (r *postRepository) Update(ctx context.Context, newPost *domains.UpdatePostDTO) (err error) {
	query := `
		UPDATE t_posts SET
			title = $1,
			content = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $3`
	_, err = r.dbpool.Exec(ctx, query, newPost.Title, newPost.Content, newPost.ID)
	return
}
func (r *postRepository) DeleteByID(ctx context.Context, postID, userID uuid.UUID) (err error) {
	query := `DELETE FROM t_posts WHERE id = $1 AND user_id = $2`
	_, err = r.dbpool.Exec(ctx, query, postID, userID)
	return
}

func (r *postRepository) GetByID(ctx context.Context, postID, userID uuid.UUID) (post *domains.InPostDTO, err error) {
	query := `SELECT u.name,u.surname,c.name, p.title, p.content, p.created_at,t.name,COUNT(l.* ) AS like_count,i.path FROM t_posts p
	INNER JOIN t_users u ON p.user_id = u.id
	INNER JOIN t_categories c ON p.category_id = c.id
	INNER JOIN t_images i ON p.id = i.post_id
	INNER JOIN t_likes l ON p.id = l.post_id
	INNER JOIN t_tags t ON p.id = t.post_id 
	WHERE ($1::uuid = uuid_nil() OR p.id = $1)  AND 
		  ($2::uuid = uuid_nil() OR p.user_id = $2)`
	err = r.dbpool.QueryRow(ctx, query, postID).Scan(&post.Name, &post.Surname, &post.Category, &post.Title, &post.Content, &post.CreatedAt, &post.Tags, &post.Likes, &post.Images)
	return
}

func (r *postRepository) GetAll(ctx context.Context, limit, offset uint64) (posts []*domains.PostDTO, err error) {
	query := `SELECT p.id,u.name,u.surname,c.name, p.title, p.content, p.created_at,t.name,COUNT(l.* ) AS like_count,COUNT(cm.*) AS cm_count ,i.path FROM t_posts p
	INNER JOIN t_users u ON p.user_id = u.id
	INNER JOIN t_categories c ON p.category_id = c.id
	INNER JOIN t_images i ON p.id = i.post_id
	INNER JOIN t_likes l ON p.id = l.post_id
	INNER JOIN t_tags t ON p.id = t.post_id
	INNER JOIN t_comments cm ON p.id = cm.post_id
	LIMIT $1 OFFSET $2`
	rows, err := r.dbpool.Query(ctx, query, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {

		post := &domains.PostDTO{}
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.UserID, &post.CreatedAt, &post.Tags, &post.Likes, &post.Comments, &post.Images)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

func (r *postRepository) GetByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset uint64) (posts []*domains.PostDTO, err error) {
	query := `SELECT p.id,u.name,u.surname,c.name, p.title, p.content, p.created_at,t.name,COUNT(l.* ) AS like_count,COUNT(cm.*) AS cm_count ,i.path FROM t_posts p
	INNER JOIN t_users u ON p.user_id = u.id
	INNER JOIN t_categories c ON p.category_id = c.id
	INNER JOIN t_images i ON p.id = i.post_id
	INNER JOIN t_likes l ON p.id = l.post_id
	INNER JOIN t_tags t ON p.id = t.post_id
	INNER JOIN t_comments cm ON p.id = cm.post_id
	WHERE c.id = $1
	LIMIT $2 OFFSET $3`
	rows, err := r.dbpool.Query(ctx, query, categoryID, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		post := &domains.PostDTO{}
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.UserID, &post.CreatedAt, &post.Tags, &post.Likes, &post.Comments, &post.Images)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return

}

func (r *postRepository) GetByTag(ctx context.Context, tagID uuid.UUID, limit, offset uint64) (posts []*domains.PostDTO, err error) {
	query := `SELECT p.id,u.name,u.surname,c.name, p.title, p.content, p.created_at,t.name,COUNT(l.* ) AS like_count,COUNT(cm.*) AS cm_count ,i.path FROM t_posts p
	INNER JOIN t_users u ON p.user_id = u.id
	INNER JOIN t_categories c ON p.category_id = c.id
	INNER JOIN t_images i ON p.id = i.post_id
	INNER JOIN t_likes l ON p.id = l.post_id
	INNER JOIN t_tags t ON p.id = t.post_id
	INNER JOIN t_comments cm ON p.id = cm.post_id
	WHERE t.id = $1
	LIMIT $2 OFFSET $3`
	rows, err := r.dbpool.Query(ctx, query, tagID, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		post := &domains.PostDTO{}
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.UserID, &post.CreatedAt, &post.Tags, &post.Likes, &post.Comments, &post.Images)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

func (r *postRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset uint64) (posts []*domains.PostDTO, err error) {
	query := `SELECT p.id,u.name,u.surname,c.name, p.title, p.content, p.created_at,t.name,COUNT(l.* ) AS like_count,COUNT(cm.*) AS cm_count ,i.path FROM t_posts p
	INNER JOIN t_users u ON p.user_id = u.id
	INNER JOIN t_categories c ON p.category_id = c.id
	INNER JOIN t_images i ON p.id = i.post_id
	INNER JOIN t_likes l ON p.id = l.post_id
	INNER JOIN t_tags t ON p.id = t.post_id
	INNER JOIN t_comments cm ON p.id = cm.post_id
	WHERE u.id = $1
	LIMIT $2 OFFSET $3`
	rows, err := r.dbpool.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		post := &domains.PostDTO{}
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.UserID, &post.CreatedAt, &post.Tags, &post.Likes, &post.Comments, &post.Images)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}
