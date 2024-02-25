package roadmap_ports

import (
	"context"

	"github.com/google/uuid"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
)

type IRoadmapService interface {
	CreateRoadmap(ctx context.Context, roadmap *roadmap_domain.Roadmap) (err error)
	UpdateRoadmap(ctx context.Context, newRoadmap *roadmap_domain.Roadmap) (err error)
	DeleteRoadmap(ctx context.Context, roadmapID uuid.UUID) (err error)
	SearchRoadmap(ctx context.Context, category_id uuid.UUID) (roadmaps []*roadmap_domain.Roadmap, err error)
}

type IPathWayService interface {
	CreatePathWay(ctx context.Context, pathway *roadmap_domain.PathWay) (err error)
	UpdatePathWay(ctx context.Context, newPathway *roadmap_domain.PathWay) (err error)
	DeletePathWay(ctx context.Context, pathWayID uuid.UUID) (err error)
	SearchPathWay(ctx context.Context, filter *roadmap_domain.PathWay) (pathways []*roadmap_domain.PathWay, err error)
}

type ICategoryService interface {
	CreateCategory(ctx context.Context, category *roadmap_domain.Category) (err error)
	UpdateCategory(ctx context.Context, newCategory *roadmap_domain.Category) (err error)
	DeleteCategory(ctx context.Context, categoryID uuid.UUID) (err error)
	SearchCategory(ctx context.Context, name string, id uuid.UUID) (categories []*roadmap_domain.Category, err error)
}

type ISubmissionService interface {
	RegisterRoadmap(ctx context.Context, submission *roadmap_domain.Submission) (err error)
	UnregisterRoadmap(ctx context.Context, submissionID uuid.UUID) (err error)
	GetRegisteredRoadmaps(ctx context.Context, filter *roadmap_domain.Submission) (submissions []*roadmap_domain.Submission, err error)
}

type IAdvanceService interface {
	AddAdvance(ctx context.Context, advance *roadmap_domain.Advance) (err error)
	DeleteAdvance(ctx context.Context, advanceID uuid.UUID) (err error)
}

type IRoadmapServices interface {
	RoadmapService() IRoadmapService
	PathWayService() IPathWayService
	SubmissionService() ISubmissionService
	AdvanceService() IAdvanceService
	CategoryService() ICategoryService
}
