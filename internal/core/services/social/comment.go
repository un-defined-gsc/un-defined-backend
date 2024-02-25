package social_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type commentService struct {
	socialRepositories social_ports.ISocialRepositories
	deps               deps_ports.IDepsServices
}

func newCommentService(
	socialRepositories social_ports.ISocialRepositories,
	deps deps_ports.IDepsServices,
) social_ports.ICommentsService {
	return &commentService{
		socialRepositories: socialRepositories,
		deps:               deps,
	}
}

func (s *commentService) CreateComment(ctx context.Context, comment *domains.CommentDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(comment); err != nil {
		return
	}
	return s.socialRepositories.CommentsRepository().Create(ctx, comment)
}

func (s *commentService) DeleteComment(ctx context.Context, commentID, userID uuid.UUID) (err error) {

	return s.socialRepositories.CommentsRepository().DeleteByID(ctx, commentID, userID)
}
