package roadmap_services

import (
	"context"

	"github.com/google/uuid"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type pathWayService struct {
	roadmapRepository roadmap_ports.IRoadmapRepositories
	deps              deps_ports.IDepsServices
}

func newPathWayService(roadmapRepository roadmap_ports.IRoadmapRepositories, deps deps_ports.IDepsServices) roadmap_ports.IPathWayService {
	return &pathWayService{
		roadmapRepository: roadmapRepository,
		deps:              deps,
	}
}

func (s *pathWayService) CreatePathWay(ctx context.Context, pathway *roadmap_domain.PathWay) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(*pathway)
	if err != nil {
		return
	}
	err = s.roadmapRepository.PathWay().Create(ctx, pathway)
	if err != nil {
		return
	}
	return
}

func (s *pathWayService) UpdatePathWay(ctx context.Context, newPathway *roadmap_domain.PathWay) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(*newPathway)
	if err != nil {
		return
	}
	err = s.roadmapRepository.PathWay().Update(ctx, newPathway)
	if err != nil {
		return
	}
	return
}

func (s *pathWayService) DeletePathWay(ctx context.Context, pathwayID uuid.UUID) (err error) {
	err = s.roadmapRepository.PathWay().Delete(ctx, pathwayID)
	if err != nil {
		return
	}
	return
}

func (s *pathWayService) SearchPathWay(ctx context.Context, filter *roadmap_domain.PathWay) (pathways []*roadmap_domain.PathWay, err error) {
	pathways, err = s.roadmapRepository.PathWay().Filter(ctx, filter)
	if err != nil {
		return
	}
	return
}
