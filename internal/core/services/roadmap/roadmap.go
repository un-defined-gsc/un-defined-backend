package roadmap_services

import (
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
)

type roadmmapService struct {
	roadmapRepository roadmap_ports.IRoadmapRepositories
	deps              deps_ports.IDepsServices
}

func NewRoadmapService(roadmapRepository roadmap_ports.IRoadmapRepositories, deps deps_ports.IDepsServices) roadmap_ports. {
	return &roadmmapService{
		roadmapRepository: roadmapRepository,
		deps:              deps,
	}
}


