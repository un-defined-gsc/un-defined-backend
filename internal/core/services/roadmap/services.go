package roadmap_services

import (
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type roadmapServices struct {
	roadmmapService   roadmap_ports.IRoadmapService
	advanceService    roadmap_ports.IAdvanceService
	submissionService roadmap_ports.ISubmissionService
	pathWayService    roadmap_ports.IPathWayService
	categoryService   roadmap_ports.ICategoryService
}

func NewRoadmapServices(roadmapRepositories roadmap_ports.IRoadmapRepositories, deps deps_ports.IDepsServices) roadmap_ports.IRoadmapServices {
	return &roadmapServices{
		roadmmapService:   newRoadmapService(roadmapRepositories, deps),
		advanceService:    newAdvanceService(roadmapRepositories, deps),
		submissionService: newSubmissionService(roadmapRepositories, deps),
		pathWayService:    newPathWayService(roadmapRepositories, deps),
		categoryService:   newCategoryService(roadmapRepositories, deps),
	}
}

func (s *roadmapServices) PathWayService() roadmap_ports.IPathWayService {
	return s.pathWayService
}

func (s *roadmapServices) RoadmapService() roadmap_ports.IRoadmapService {
	return s.roadmmapService
}

func (s *roadmapServices) AdvanceService() roadmap_ports.IAdvanceService {
	return s.advanceService
}

func (s *roadmapServices) SubmissionService() roadmap_ports.ISubmissionService {
	return s.submissionService
}

func (s *roadmapServices) CategoryService() roadmap_ports.ICategoryService {
	return s.categoryService
}
