package social_ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
)

type IPostsService interface {
	CreatePost(ctx context.Context, post *domains.CratePostDTO) (err error)
	UpdatePost(ctx context.Context, newPost *domains.UpdatePostDTO) (err error)
	DeletePost(ctx context.Context, postID uuid.UUID, userID uuid.UUID) (err error)
	GetPost(ctx context.Context, postID uuid.UUID, userID uuid.UUID) (post *domain.InPostDTO, err error)
	GetPosts(ctx context.Context, limit, offset uint64) (posts []*domains.PostDTO, err error)
	GetPostByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset uint64) (posts []*domains.PostDTO, err error)
	GetPostByTag(ctx context.Context, tagID uuid.UUID, limit, offset uint64) (posts []*domains.PostDTO, err error)
	GetPostByUserID(ctx context.Context, userID uuid.UUID, limit, offset uint64) (posts []*domains.PostDTO, err error)
}

type ICategoriesService interface {
	GetCategories(ctx context.Context) (categories []*domains.CategoryDTO, err error)
}

type ICommentsService interface {
	CreateComment(ctx context.Context, comment *social_domain.Comment) (err error)
	UpdateComment(ctx context.Context, newComment *social_domain.Comment) (err error)
	DeleteComment(ctx context.Context, commentID uuid.UUID) (err error)
	GetComment(ctx context.Context, commentID uuid.UUID) (comment *social_domain.Comment, err error)
	GetComments(ctx context.Context, postID uuid.UUID, limit, offset uint64) (comments []*social_domain.Comment, err error)
}

type IImagesService interface {
	UploadImage(ctx context.Context, image *social_domain.Image) (err error)
	DeleteImage(ctx context.Context, imageID uuid.UUID) (err error)
	UpdateImage(ctx context.Context, newImage *social_domain.Image) (err error)
}

type ILikesService interface {
	Like(ctx context.Context, like *social_domain.Like) (err error)
	UnLike(ctx context.Context, likeID uuid.UUID) (err error)
}

type ISocialServices interface {
	PostsService() IPostsService
	CategoriesService() ICategoriesService
	CommentsService() ICommentsService

	ImagesService() IImagesService
}
