package social_service

import (
	"context"

	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
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

func (s *likesService) Like(ctx context.Context, like *domains.LikeDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(like); err != nil {
		return
	}
	return s.socialRepositories.LikesRepository().Like(ctx, like)
}

func (s *likesService) UnLike(ctx context.Context, like *domains.LikeDTO) (err error) {
	return s.socialRepositories.LikesRepository().UnLike(ctx, like)
}
