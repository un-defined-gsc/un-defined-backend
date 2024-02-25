package social_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type imagesRepository struct {
	dbpool *pgxpool.Pool
}

func NewImagesRepository(dbpool *pgxpool.Pool) social_ports.IImagesRepository {
	return &imagesRepository{
		dbpool: dbpool,
	}
}

func (r *imagesRepository) Create(ctx context.Context, image *social_domain.Image) (err error) {
	_, err = r.dbpool.Exec(ctx, `INSERT INTO t_images (url, user_id,post_id,category) VALUES ($1, $2,$3,$4)`, image.Path, image.UserID, image.PostID, image.Category)
	return
}

func (r *imagesRepository) Update(ctx context.Context, newImage *social_domain.Image) (err error) {
	_, err = r.dbpool.Exec(ctx, `
		UPDATE t_images SET url = $1 WHERE id = $2
	`, newImage.Path, newImage.ID)
	return
}

func (r *imagesRepository) DeleteByID(ctx context.Context, imageID uuid.UUID) (err error) {
	_, err = r.dbpool.Exec(ctx, `
		DELETE FROM t_images WHERE id = $1
	`, imageID)
	return
}

func (r *imagesRepository) GetByPath(ctx context.Context, imagePath string) (imageID uuid.UUID, err error) {
	err = r.dbpool.QueryRow(ctx, `
		SELECT id FROM t_images WHERE url = $1
	`, imagePath).Scan(&imageID)
	return

}
