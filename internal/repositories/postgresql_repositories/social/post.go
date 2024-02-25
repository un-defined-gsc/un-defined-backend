package social_repositories

import (
	"context"
	"errors"
	"strings"

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

func (r *postRepository) Create(ctx context.Context, post *domains.CratePostDTO, categoryID uuid.UUID) (err error) {
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
				return service_errors.ErrDataDuplication
			}
		}
	}
	return
}
func (r *postRepository) Update(ctx context.Context, newPost *domains.UpdatePostDTO) (err error) {
	query := `
		UPDATE t_posts SET
			title = $1,
			content = $2
		WHERE id = $3`
	_, err = r.dbpool.Exec(ctx, query, newPost.Title, newPost.Content, newPost.ID)
	return
}
func (r *postRepository) DeleteByID(ctx context.Context, postID, userID uuid.UUID) (err error) {
	query := `DELETE FROM t_posts WHERE id = $1 AND user_id = $2`
	_, err = r.dbpool.Exec(ctx, query, postID, userID)
	return
}

func (r *postRepository) GetByID(ctx context.Context, postID, userID uuid.UUID) (post domains.InPostDTO, err error) {
	query := `
	SELECT  u.first_name, u.last_name, c.name, p.title, p.content, p.created_at, 
		string_agg(t.name, ', ') AS tags,
		COUNT(DISTINCT l.id) AS like_count, 
		COUNT(DISTINCT cm.id) AS cm_count, 
		array_agg(i.url) AS images
	FROM t_posts p
		INNER JOIN t_users u ON p.user_id = u.id
		INNER JOIN t_categories c ON p.category_id = c.id
		INNER JOIN t_images i ON p.id = i.post_id
		LEFT JOIN t_likes l ON p.id = l.post_id
		LEFT JOIN t_tags t ON p.id = t.post_id
		LEFT JOIN t_comments cm ON p.id = cm.post_id
        WHERE ($1::uuid = uuid_nil() OR p.id = $1)  AND 
              ($2::uuid = uuid_nil() OR p.user_id = $2)
			  GROUP BY p.id, u.first_name, u.last_name, c.name, p.title, p.content, p.created_at
    `
	var tags string
	var images []string
	var commentcount int

	err = r.dbpool.QueryRow(ctx, query, postID, userID).Scan(&post.Name, &post.Surname, &post.Category, &post.Title, &post.Content, &post.CreatedAt, &tags, &post.Likes, &commentcount, &images)
	if err != nil {
		return
	}

	tagNames := strings.Split(tags, ", ")
	for _, tagName := range tagNames {
		post.Tags = append(post.Tags, domains.TagDTO{Name: tagName})
	}

	post.Images = images

	return
}

func (r *postRepository) GetAll(ctx context.Context, limit, offset uint64) (posts []*domains.PostDTO, err error) {
	query := `SELECT p.id, u.first_name, u.last_name, c.name, p.title, p.content, p.created_at, 
           string_agg(t.name, ', ') AS tags,
           COUNT(DISTINCT l.id) AS like_count, 
           COUNT(DISTINCT cm.id) AS cm_count, 
           array_agg(i.url) AS images
	FROM t_posts p
	INNER JOIN t_users u ON p.user_id = u.id
	INNER JOIN t_categories c ON p.category_id = c.id
	INNER JOIN t_images i ON p.id = i.post_id
	LEFT JOIN t_likes l ON p.id = l.post_id
	LEFT JOIN t_tags t ON p.id = t.post_id
	LEFT JOIN t_comments cm ON p.id = cm.post_id
	GROUP BY p.id, u.first_name, u.last_name, c.name, p.title, p.content, p.created_at
	LIMIT $1 OFFSET $2;
`

	rows, err := r.dbpool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &domains.PostDTO{}
		var tags string
		var images []string
		err := rows.Scan(&post.ID, &post.Name, &post.Surname, &post.Category, &post.Title, &post.Content, &post.CreatedAt, &tags, &post.Likes, &post.Comments, &images)
		if err != nil {
			return nil, err
		}
		tagNames := strings.Split(tags, ", ")
		for _, tagName := range tagNames {
			post.Tags = append(post.Tags, domains.TagDTO{Name: tagName})
		}
		post.Images = images
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *postRepository) GetByPostFilter(ctx context.Context, categoryID, userID uuid.UUID, tag string, limit, offset uint64) (posts []*domains.PostDTO, err error) {
	query := `SELECT p.id,u.first_name,u.last_name,c.name, p.title, p.content, p.created_at,
	string_agg(t.name, ', ') AS tags,
	COUNT(l.* ) AS like_count,
	COUNT(cm.*) AS cm_count ,
	array_agg(i.url) AS images
	FROM t_posts p
	INNER JOIN t_users u ON p.user_id = u.id
	INNER JOIN t_categories c ON p.category_id = c.id
	INNER JOIN t_images i ON p.id = i.post_id
	LEFT JOIN t_likes l ON p.id = l.post_id
	LEFT JOIN t_tags t ON p.id = t.post_id
	LEFT JOIN t_comments cm ON p.id = cm.post_id
	WHERE ($1::uuid = uuid_nil() OR c.id = $1) AND
		  ($2::uuid = uuid_nil() OR p.user_id = $2) AND
		  ($3::text = '' OR t.name = $3)
	GROUP BY p.id, u.first_name, u.last_name, c.name, p.title, p.content, p.created_at,t.name,i.url
	LIMIT $4 OFFSET $5`
	rows, err := r.dbpool.Query(ctx, query, categoryID, userID, tag, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		post := &domains.PostDTO{}
		var tags string
		var images []string
		err = rows.Scan(&post.ID, &post.Name, &post.Surname, &post.Category, &post.Title, &post.Content, &post.CreatedAt, &tags, &post.Likes, &post.Comments, &images)
		if err != nil {
			return
		}
		tagNames := strings.Split(tags, ", ")
		for _, tagName := range tagNames {
			post.Tags = append(post.Tags, domains.TagDTO{Name: tagName})
		}
		post.Images = images
		posts = append(posts, post)
	}
	return

}
