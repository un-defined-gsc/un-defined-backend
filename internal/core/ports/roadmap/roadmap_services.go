package roadmap_ports

import (
	"context"

	"github.com/google/uuid"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
)

type IRoadmapServices interface {
	CreateRoadmap(ctx context.Context, roadmap *roadmap_domain.Roadmap) (err error)
	UpdateRoadmap(ctx context.Context, newRoadmap *roadmap_domain.Roadmap) (err error)
	DeleteRoadmap(ctx context.Context, roadmapID uuid.UUID) (err error)
	SearchRoadmap(ctx context.Context, filter *roadmap_domain.Roadmap) (roadmaps []*roadmap_domain.Roadmap, err error)
}

type IPathWayServices interface {
	CreatePathWay(ctx context.Context, pathway *roadmap_domain.PathWay) (err error)
	UpdatePathWay(ctx context.Context, newPathway *roadmap_domain.PathWay) (err error)
	DeletePathWay(ctx context.Context, pathWayID uuid.UUID) (err error)
	SearchPathWay(ctx context.Context, filter *roadmap_domain.PathWay) (pathways []*roadmap_domain.PathWay, err error)
}

type ISubmissionServices interface {
	RegisterRoadmap(ctx context.Context, submission *roadmap_domain.Submission) (err error)
	UnregisterRoadmap(ctx context.Context, submissionID uuid.UUID) (err error)
	GetRegisteredRoadmaps(ctx context.Context, filter *roadmap_domain.Submission) (submissions []*roadmap_domain.Submission, err error)
}

type IAdvanceServices interface {
	AddAdvance(ctx context.Context, advance *roadmap_domain.Advance) (err error)
	DeleteAdvance(ctx context.Context, advanceID uuid.UUID) (err error)
}
