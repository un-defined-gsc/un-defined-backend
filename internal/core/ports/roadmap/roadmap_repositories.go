package roadmap_ports

import (
	"context"

	"github.com/google/uuid"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
)

type IRoadmapRepository interface {
	Create(ctx context.Context, roadmap *roadmap_domain.Roadmap) (err error)
	Update(ctx context.Context, newRoadmap *roadmap_domain.Roadmap) (err error)
	Delete(ctx context.Context, roadmapID uuid.UUID) (err error)
	Filter(ctx context.Context, filter *roadmap_domain.Roadmap) (roadmaps []*roadmap_domain.Roadmap, err error)
}

type IPathWayRepository interface {
	Create(ctx context.Context, pathway *roadmap_domain.PathWay) (err error)
	Update(ctx context.Context, newPathway *roadmap_domain.PathWay) (err error)
	Delete(ctx context.Context, pathWayID uuid.UUID) (err error)
	Filter(ctx context.Context, filter *roadmap_domain.PathWay) (pathways []*roadmap_domain.PathWay, err error)
}

type ISubmissionRepository interface {
	Create(ctx context.Context, submission *roadmap_domain.Submission) (err error)
	Delete(ctx context.Context, submissionID uuid.UUID) (err error)
	Filter(ctx context.Context, filter *roadmap_domain.Submission) (submissions []*roadmap_domain.Submission, err error)
}

type IAdvanceRepository interface {
	Create(ctx context.Context, advance *roadmap_domain.Advance) (err error)
	Delete(ctx context.Context, advanceID uuid.UUID) (err error)
	Filter(ctx context.Context, filter *roadmap_domain.Advance) (advances []*roadmap_domain.Advance, err error)
}

type ICategoryRepository interface {
	Create(ctx context.Context, category *roadmap_domain.Category) (err error)
	Update(ctx context.Context, newCategory *roadmap_domain.Category) (err error)
	Delete(ctx context.Context, categoryID uuid.UUID) (err error)
	Filter(ctx context.Context, filter *roadmap_domain.Category) (categories []*roadmap_domain.Category, err error)
}

type IRoadmapRepositories interface {
	Roadmap() IRoadmapRepository
	PathWay() IPathWayRepository
	Submission() ISubmissionRepository
	Advance() IAdvanceRepository
	Category() ICategoryRepository
}
