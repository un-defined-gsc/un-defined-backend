package roadmap_services

import (
	"context"

	"github.com/google/uuid"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type roadmmapService struct {
	roadmapRepository roadmap_ports.IRoadmapRepositories
	deps              deps_ports.IDepsServices
}

func newRoadmapService(roadmapRepository roadmap_ports.IRoadmapRepositories, deps deps_ports.IDepsServices) roadmap_ports.IRoadmapService {
	return &roadmmapService{
		roadmapRepository: roadmapRepository,
		deps:              deps,
	}
}
func (s *roadmmapService) CreateRoadmap(ctx context.Context, roadmap *roadmap_domain.Roadmap) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(*roadmap)
	if err != nil {
		return
	}
	err = s.roadmapRepository.Roadmap().Create(ctx, roadmap)
	if err != nil {
		return
	}
	return
}

func (s *roadmmapService) UpdateRoadmap(ctx context.Context, newRoadmap *roadmap_domain.Roadmap) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(*newRoadmap)
	if err != nil {
		return
	}
	err = s.roadmapRepository.Roadmap().Update(ctx, newRoadmap)
	if err != nil {
		return
	}
	return
}

func (s *roadmmapService) DeleteRoadmap(ctx context.Context, roadmapID uuid.UUID) (err error) {
	err = s.roadmapRepository.Roadmap().Delete(ctx, roadmapID)
	if err != nil {
		return
	}
	return
}

func (s *roadmmapService) SearchRoadmap(ctx context.Context, categor_id uuid.UUID) (roadmaps []*roadmap_domain.Roadmap, err error) {
	roadmaps, err = s.roadmapRepository.Roadmap().Filter(ctx, &roadmap_domain.Roadmap{CategoryID: categor_id})
	if err != nil {
		return
	}
	return
}
