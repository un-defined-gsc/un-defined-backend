package social_service

import (
	"context"

	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
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

	l, err := s.socialRepositories.LikesRepository().GetLikesByPostID(ctx, like)
	if len(l) > 0 {
		return service_errors.ErrLikedAlreadyExists
	}
	if err != nil {
		return
	}
	_, err = s.socialRepositories.LikesRepository().Like(ctx, like)
	return

}

func (s *likesService) UnLike(ctx context.Context, like *domains.LikeDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(like); err != nil {
		return
	}
	l, err := s.socialRepositories.LikesRepository().GetLikesByPostID(ctx, like)
	if len(l) == 0 {
		return service_errors.ErrLikeNotFound
	}
	if err != nil {
		return
	}
	_, err = s.socialRepositories.LikesRepository().UnLike(ctx, like)
	return
}
