package social_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type CategoryRepository struct {
	dbpool *pgxpool.Pool
}

func NewCategoryRepository(dbpool *pgxpool.Pool) social_ports.ICategoriesRepository {
	return &CategoryRepository{
		dbpool: dbpool,
	}
}

func (r *CategoryRepository) GetAll(ctx context.Context) (categories []*domains.CategoryDTO, err error) {
	query := `SELECT id, name FROM t_categories`
	rows, err := r.dbpool.Query(ctx, query)
	if err != nil {
		return
	}
	for rows.Next() {
		var category domains.CategoryDTO
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return
		}
		categories = append(categories, &category)
	}
	return
}

func (r *CategoryRepository) GetByName(ctx context.Context, name string) (categoryID uuid.UUID, err error) {
	query := `SELECT id FROM t_categories WHERE name = $1`
	err = r.dbpool.QueryRow(ctx, query, name).Scan(&categoryID)
	return
}
