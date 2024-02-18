package social_ports

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
)

type IPostsService interface {
	CreatePost(ctx context.Context, post *social_domain.Post) (err error)
	UpdatePost(ctx context.Context, newPost *social_domain.Post) (err error)
	DeletePost(ctx context.Context, postID uuid.UUID) (err error)
	GetPost(ctx context.Context, postID uuid.UUID) (post *domain.InPostDTO, err error)
	GetPosts(ctx context.Context, limit, offset uint64) (posts []*social_domain.Post, err error)
	GetPostByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error)
	GetPostByTag(ctx context.Context, tagID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error)
	GetPostByUserID(ctx context.Context, userID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error)
}

type ICategoriesService interface {
	CreateCategory(ctx context.Context, category *social_domain.Category) (err error)
	DeleteCategory(ctx context.Context, categoryID uuid.UUID) (err error)
	GetCategory(ctx context.Context, categoryID uuid.UUID) (category *social_domain.Category, err error)
	GetCategories(ctx context.Context, limit, offset uint64) (categories []*social_domain.Category, err error)
}

type ICommentsService interface {
	CreateComment(ctx context.Context, comment *social_domain.Comment) (err error)
	UpdateComment(ctx context.Context, newComment *social_domain.Comment) (err error)
	DeleteComment(ctx context.Context, commentID uuid.UUID) (err error)
	GetComment(ctx context.Context, commentID uuid.UUID) (comment *social_domain.Comment, err error)
	GetComments(ctx context.Context, postID uuid.UUID, limit, offset uint64) (comments []*social_domain.Comment, err error)
}

type ITagsService interface {
	CreateTag(ctx context.Context, tag *social_domain.Tag) (err error)
	DeleteTag(ctx context.Context, tagID uuid.UUID) (err error)
	GetTag(ctx context.Context, tagID uuid.UUID) (tag *social_domain.Tag, err error)
	GetTags(ctx context.Context, limit, offset uint64) (tags []*social_domain.Tag, err error)
}

type ISocialServices interface {
	PostsService() IPostsService
	CategoriesService() ICategoriesService
	CommentsService() ICommentsService
	TagsService() ITagsService
}
