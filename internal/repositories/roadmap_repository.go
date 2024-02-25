package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
	roadmap_ps_repositories "github.com/un-defined-gsc/un-defined-backend/internal/repositories/postgresql_repositories/roadmap"
)

type roadmapRepositories struct {
	roadmapRepository    roadmap_ports.IRoadmapRepository
	advanceRepository    roadmap_ports.IAdvanceRepository
	submissionRepository roadmap_ports.ISubmissionRepository
	pathWayRepository    roadmap_ports.IPathWayRepository
	categoryRepository   roadmap_ports.ICategoryRepository
}

func NewRoadmapRepositories(dbpool *pgxpool.Pool) roadmap_ports.IRoadmapRepositories {
	return &roadmapRepositories{
		roadmapRepository:    roadmap_ps_repositories.NewRoadmapRepository(dbpool),
		advanceRepository:    roadmap_ps_repositories.NewAdvanceRepository(dbpool),
		submissionRepository: roadmap_ps_repositories.NewSubmissionRepository(dbpool),
		pathWayRepository:    roadmap_ps_repositories.NewPathWayRepository(dbpool),
		categoryRepository:   roadmap_ps_repositories.NewCategoryRepository(dbpool),
	}
}

func (r *roadmapRepositories) Roadmap() roadmap_ports.IRoadmapRepository {
	return r.roadmapRepository
}

func (r *roadmapRepositories) Advance() roadmap_ports.IAdvanceRepository {
	return r.advanceRepository
}

func (r *roadmapRepositories) Submission() roadmap_ports.ISubmissionRepository {
	return r.submissionRepository
}

func (r *roadmapRepositories) PathWay() roadmap_ports.IPathWayRepository {
	return r.pathWayRepository
}

func (r *roadmapRepositories) Category() roadmap_ports.ICategoryRepository {
	return r.categoryRepository
}
