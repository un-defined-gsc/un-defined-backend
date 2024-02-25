package roadmap_services

import (
	"context"

	"github.com/google/uuid"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type submissionService struct {
	roadmapRepository roadmap_ports.IRoadmapRepositories
	deps              deps_ports.IDepsServices
}

func newSubmissionService(roadmapRepository roadmap_ports.IRoadmapRepositories, deps deps_ports.IDepsServices) roadmap_ports.ISubmissionService {
	return &submissionService{
		roadmapRepository: roadmapRepository,
		deps:              deps,
	}
}

func (s *submissionService) RegisterRoadmap(ctx context.Context, submission *roadmap_domain.Submission) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(*submission)
	if err != nil {
		return
	}
	err = s.roadmapRepository.Submission().Create(ctx, submission)
	if err != nil {
		return
	}
	return
}

func (s *submissionService) UnregisterRoadmap(ctx context.Context, submissionID uuid.UUID) (err error) {
	err = s.roadmapRepository.Submission().Delete(ctx, submissionID)
	if err != nil {
		return
	}
	return
}

func (s *submissionService) GetRegisteredRoadmaps(ctx context.Context, filter *roadmap_domain.Submission) (submissions []*roadmap_domain.Submission, err error) {
	return s.roadmapRepository.Submission().Filter(ctx, filter)
}
