package social_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type postService struct {
	socialRepositories social_ports.ISocialRepositories
	deps               deps_ports.IDepsServices
}

func newPostService(
	socialRepositories social_ports.ISocialRepositories,
	deps deps_ports.IDepsServices,
) social_ports.IPostsService {
	return &postService{
		socialRepositories: socialRepositories,
		deps:               deps,
	}
}

func (s *postService) CreatePost(ctx context.Context, post *domains.CratePostDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(post); err != nil {
		return
	}
	for idx := range post.Tags {
		err := s.deps.CensorService().CensorText(&post.Tags[idx])
		if err != nil {
			return err
		}
	}

	err = s.deps.CensorService().CensorText(&post.Title, &post.Content)
	if err != nil {
		return
	}
	return s.socialRepositories.PostsRepository().Create(ctx, post)
}

func (s *postService) UpdatePost(ctx context.Context, newPost *domains.UpdatePostDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(newPost); err != nil {
		return
	}

	if err != nil {
		return
	}
	_, err = s.socialRepositories.PostsRepository().GetByID(ctx, newPost.ID)
	if err != nil {
		return
	}

	for idx := range newPost.Tags {
		err := s.deps.CensorService().CensorText(&newPost.Tags[idx])
		if err != nil {
			return err
		}
	}

	err = s.deps.CensorService().CensorText(&newPost.Title, &newPost.Content)
	if err != nil {
		return
	}

	return s.socialRepositories.PostsRepository().Update(ctx, newPost)
}

func (s *postService) DeletePost(ctx context.Context, postID uuid.UUID, userID uuid.UUID) (err error) {

	_, err = s.socialRepositories.PostsRepository().GetByUserIDAndPostID(ctx, userID, postID)
	if err != nil {
		return
	}

	return s.socialRepositories.PostsRepository().DeleteByID(ctx, postID, userID)
}

func (s *postService) GetPost(ctx context.Context, postID uuid.UUID) (post *domain.InPostDTO, err error) {

	return s.socialRepositories.PostsRepository().GetByID(ctx, postID)
}

func (s *postService) GetPosts(ctx context.Context, limit, offset uint64) (posts []*social_domain.Post, err error) {

	return s.socialRepositories.PostsRepository().GetAll(ctx, limit, offset)
}
func (s *postService) GetPostByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error) {

	return s.socialRepositories.PostsRepository().GetByCategory(ctx, categoryID, limit, offset)
}

func (s *postService) GetPostByTag(ctx context.Context, tagID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error) {

	return s.socialRepositories.PostsRepository().GetByTag(ctx, tagID, limit, offset)
}

func (s *postService) GetPostByUserID(ctx context.Context, userID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error) {

	return s.socialRepositories.PostsRepository().GetByUserID(ctx, userID, limit, offset)
}
