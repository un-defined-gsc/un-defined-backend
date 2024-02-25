package roadmap_services

import (
	"context"

	"github.com/google/uuid"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type categoryService struct {
	roadmapRepositories roadmap_ports.IRoadmapRepositories
	deps                deps_ports.IDepsServices
}

func newCategoryService(roadmapRepositories roadmap_ports.IRoadmapRepositories, deps deps_ports.IDepsServices) roadmap_ports.ICategoryService {
	return &categoryService{
		roadmapRepositories: roadmapRepositories,
		deps:                deps,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, category *roadmap_domain.Category) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(*category)
	if err != nil {
		return
	}
	err = s.roadmapRepositories.Category().Create(ctx, category)
	if err != nil {
		return
	}
	return
}

func (s *categoryService) UpdateCategory(ctx context.Context, newCategory *roadmap_domain.Category) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(*newCategory)
	if err != nil {
		return
	}
	err = s.roadmapRepositories.Category().Update(ctx, newCategory)
	if err != nil {
		return
	}
	return
}

func (s *categoryService) DeleteCategory(ctx context.Context, categoryID uuid.UUID) (err error) {
	err = s.roadmapRepositories.Category().Delete(ctx, categoryID)
	if err != nil {
		return
	}
	return
}

func (s *categoryService) SearchCategory(ctx context.Context, name string, categoryId uuid.UUID) (categories []*roadmap_domain.Category, err error) {
	filter := &roadmap_domain.Category{
		Base: base_domain.Base{
			ID: categoryId,
		},
		Name: name,
	}
	categories, err = s.roadmapRepositories.Category().Filter(ctx, filter)
	if err != nil {
		return
	}
	return
}
