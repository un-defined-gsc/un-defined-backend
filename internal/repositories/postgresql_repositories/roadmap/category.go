package roadmap_ps_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type categoryRepository struct {
	dbpool *pgxpool.Pool
}

func NewCategoryRepository(dbpool *pgxpool.Pool) roadmap_ports.ICategoryRepository {
	return &categoryRepository{
		dbpool: dbpool,
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *roadmap_domain.Category) (err error) {
	query := `INSERT INTO 
	t_categories 
	(name) VALUES ($1) returning id`
	err = r.dbpool.QueryRow(ctx, query, category.Name).Scan(&category.ID)
	return

}

func (r *categoryRepository) Update(ctx context.Context, newCategory *roadmap_domain.Category) (err error) {
	query := `UPDATE t_categories SET name = $1 WHERE id = $2`
	_, err = r.dbpool.Exec(ctx, query, newCategory.Name, newCategory.ID)
	return

}

func (r *categoryRepository) Delete(ctx context.Context, categoryID uuid.UUID) (err error) {
	query := `DELETE FROM t_categories WHERE id = $1`
	_, err = r.dbpool.Exec(ctx, query, categoryID)
	return

}

func (r *categoryRepository) Filter(ctx context.Context, filter *roadmap_domain.Category) (categories []*roadmap_domain.Category, err error) {
	query := `SELECT * FROM t_categories WHERE
	($1::uuid = uuid_nil() OR id = $1) AND
	($2::text = "" OR name ILIKE $2 || '%')`
	rows, err := r.dbpool.Query(ctx, query, filter.ID, filter.Name)
	if err != nil {
		return
	}
	defer rows.Close()
	categories, err = pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[roadmap_domain.Category])
	return
}
