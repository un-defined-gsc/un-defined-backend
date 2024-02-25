package roadmap_services

import (
	"context"

	"github.com/google/uuid"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type advanceService struct {
	roadmapRepositories roadmap_ports.IRoadmapRepositories
	deps                deps_ports.IDepsServices
}

func newAdvanceService(roadmapRepositories roadmap_ports.IRoadmapRepositories, deps deps_ports.IDepsServices) roadmap_ports.IAdvanceService {
	return &advanceService{
		roadmapRepositories: roadmapRepositories,
		deps:                deps,
	}
}

func (s *advanceService) AddAdvance(ctx context.Context, advance *roadmap_domain.Advance) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(*advance)
	if err != nil {
		return
	}
	err = s.roadmapRepositories.Advance().Create(ctx, advance)
	if err != nil {
		return
	}
	return
}

func (s *advanceService) DeleteAdvance(ctx context.Context, advanceID uuid.UUID) (err error) {
	err = s.roadmapRepositories.Advance().Delete(ctx, advanceID)
	if err != nil {
		return
	}
	return
}
