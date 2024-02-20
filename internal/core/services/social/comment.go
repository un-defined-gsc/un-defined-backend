package social_service

import (
	"context"

	"github.com/google/uuid"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
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

func (s *commentService) CreateComment(ctx context.Context, comment *social_domain.Comment) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(comment); err != nil {
		return
	}
	err = s.deps.CensorService().CensorText(&comment.Body)
	if err != nil {
		return
	}
	return s.socialRepositories.CommentsRepository().Create(ctx, comment)
}

func (s *commentService) UpdateComment(ctx context.Context, newComment *social_domain.Comment) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(newComment); err != nil {
		return
	}
	_, err = s.socialRepositories.CommentsRepository().GetByUUID(ctx, *newComment.ID)
	if err != nil {
		return
	}
	err = s.deps.CensorService().CensorText(&newComment.Body)
	if err != nil {
		return
	}
	return s.socialRepositories.CommentsRepository().Update(ctx, newComment)
}

func (s *commentService) DeleteComment(ctx context.Context, commentID uuid.UUID) (err error) {

	return s.socialRepositories.CommentsRepository().DeleteByID(ctx, commentID)
}

func (s *commentService) GetComment(ctx context.Context, commentID uuid.UUID) (comment *social_domain.Comment, err error) {
	return s.socialRepositories.CommentsRepository().GetByUUID(ctx, commentID)
}

func (s *commentService) GetComments(ctx context.Context, postID uuid.UUID, limit, offset uint64) (comments []*social_domain.Comment, err error) {
	return s.socialRepositories.CommentsRepository().GetAllByPostID(ctx, postID, limit, offset)
}
