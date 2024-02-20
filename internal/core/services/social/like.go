package social_service

import (
	"context"

	"github.com/google/uuid"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type likesService struct {
	socialRepositories social_ports.ISocialRepositories
	deps               deps_ports.IDepsServices
}

func newLikesService(
	socialRepositories social_ports.ISocialRepositories,
	deps deps_ports.IDepsServices,
) social_ports.ILikesService {
	return &likesService{
		socialRepositories: socialRepositories,
		deps:               deps,
	}
}

func (s *likesService) Like(ctx context.Context, like *social_domain.Like) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(like); err != nil {
		return
	}
	return s.socialRepositories.LikesRepository().Like(ctx, like)
}

func (s *likesService) UnLike(ctx context.Context, likeID uuid.UUID) (err error) {
	return s.socialRepositories.LikesRepository().UnLikeByID(ctx, likeID)
}
