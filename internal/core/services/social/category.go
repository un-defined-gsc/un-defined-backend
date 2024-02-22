package social_service

import (
	"context"

	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
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

func (s *categoryService) GetCategories(ctx context.Context) (categories []*domains.CategoryDTO, err error) {

	return s.socialRepositories.CategoriesRepository().GetAll(ctx)
}
