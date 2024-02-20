package social_service

import (
	"context"

	"github.com/google/uuid"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type categoryService struct {
	socialRepositories social_ports.ISocialRepositories
	deps               deps_ports.IDepsServices
}

func newCategoriesService(
	socialRepositories social_ports.ISocialRepositories,
	deps deps_ports.IDepsServices,
) social_ports.ICategoriesService {
	return &categoryService{
		socialRepositories: socialRepositories,
		deps:               deps,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, category *social_domain.Category) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(category); err != nil {
		return
	}
	err = s.deps.CensorService().CensorText(&category.Name)
	if err != nil {
		return
	}
	return s.socialRepositories.CategoriesRepository().Create(ctx, category)
}

func (s *categoryService) DeleteCategory(ctx context.Context, categoryID uuid.UUID) (err error) {
	return s.socialRepositories.CategoriesRepository().DeleteByID(ctx, categoryID)
}

func (s *categoryService) GetCategory(ctx context.Context, categoryID uuid.UUID) (category *social_domain.Category, err error) {
	return s.socialRepositories.CategoriesRepository().GetByID(ctx, categoryID)
}

func (s *categoryService) GetCategories(ctx context.Context, limit, offsett uint64) (categories []*social_domain.Category, err error) {

	return s.socialRepositories.CategoriesRepository().GetAll(ctx, limit, offsett)
}
