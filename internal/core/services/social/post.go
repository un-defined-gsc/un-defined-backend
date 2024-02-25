package social_service

import (
	"context"
	"time"

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
	id, err := s.socialRepositories.CategoriesRepository().GetByName(ctx, post.Category)
	if err != nil {
		return
	}
	err = s.socialRepositories.PostsRepository().Create(ctx, post, id)
	if err != nil {
		return
	}
	for _, tag := range post.Tags {
		err = s.socialRepositories.TagsRepository().Create(ctx, &domains.CrateTagDTO{
			Name:   tag,
			UserID: post.UserID,
			PostID: post.ID,
		})
		if err != nil {
			return err
		}
	}
	for _, image := range post.Image {
		err = s.socialRepositories.ImagesRepository().Create(ctx, &social_domain.Image{
			UserID:   post.UserID,
			PostID:   post.ID,
			Path:     image,
			Category: "post",
		})
		if err != nil {
			return err
		}
	}

	return
}

func (s *postService) UpdatePost(ctx context.Context, newPost *domains.UpdatePostDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(newPost); err != nil {
		return
	}
	if err != nil {
		return
	}

	_, err = s.socialRepositories.PostsRepository().GetByID(ctx, newPost.ID, newPost.UserID)

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

	_, err = s.socialRepositories.PostsRepository().GetByID(ctx, postID, userID)
	if err != nil {
		return
	}

	return s.socialRepositories.PostsRepository().DeleteByID(ctx, postID, userID)
}

func (s *postService) GetPost(ctx context.Context, postID, userID uuid.UUID, limit, offset uint64) (post *domain.InPostDTO, err error) {

	newPost, err := s.socialRepositories.PostsRepository().GetByID(ctx, postID, userID)
	if err != nil {
		return
	}
	comment, err := s.socialRepositories.CommentsRepository().GetAllByPostID(ctx, postID, limit, offset)
	if err != nil {
		return
	}
	newPost.Comments = comment
	newPost.Editable = true
	newPost.Deleteable = true
	// son 24 saat kontrolü yapılacak
	if newPost.CreatedAt.AddDate(0, 0, 1).After(time.Now()) {
		// burada hata dönmesi gerekiyor
		newPost.Editable = false
	}

	return &newPost, nil

}

func (s *postService) GetPosts(ctx context.Context, limit, offset uint64) (posts []*domains.PostDTO, err error) {

	return s.socialRepositories.PostsRepository().GetAll(ctx, limit, offset)
}

func (s *postService) GetPostByFilter(ctx context.Context, categoryID, userID uuid.UUID, tag string, limit, offset uint64) (posts []*domains.PostDTO, err error) {

	return s.socialRepositories.PostsRepository().GetByPostFilter(ctx, categoryID, userID, tag, limit, offset)
}
