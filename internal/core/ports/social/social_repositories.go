package social_ports

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
)

type IPostsRepository interface {

	// Post Table Commands //
	Create(ctx context.Context, post *social_domain.Post) (err error)
	Update(ctx context.Context, newPost *social_domain.Post) (err error)
	DeleteByID(ctx context.Context, postID uuid.UUID) (err error)

	// End Post Table Commands //

	// Post Table Queries //
	GetByID(ctx context.Context, postID uuid.UUID) (post *domain.InPostDTO, err error)
	GetAll(ctx context.Context, limit, offset uint64) (posts []*social_domain.Post, err error)
	GetByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error)
	GetByTag(ctx context.Context, tagID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset uint64) (posts []*social_domain.Post, err error)
	GetByUserIDAndPostID(ctx context.Context, userID, postID uuid.UUID) (post *social_domain.Post, err error)

	// End Post Table Queries //

}

type ICategoriesRepository interface {

	// Category Table Commands //
	Create(ctx context.Context, category *social_domain.Category) (err error)
	Update(ctx context.Context, newCategory *social_domain.Category) (err error)
	DeleteByID(ctx context.Context, categoryID uuid.UUID) (err error)
	// End Category Table Commands //

	// Category Table Queries //

	GetByID(ctx context.Context, categoryID uuid.UUID) (category *social_domain.Category, err error)
	GetAll(ctx context.Context, limit, offset uint64) (categories []*social_domain.Category, err error)
	// End Category Table Queries //

}

type ICommentsRepository interface {

	// Comment Table Commands //
	Create(ctx context.Context, comment *social_domain.Comment) (err error)
	Update(ctx context.Context, newComment *social_domain.Comment) (err error)
	DeleteByID(ctx context.Context, commentID uuid.UUID) (err error)
	// End Comment Table Commands //

	// Comment Table Queries //
	GetByUUID(ctx context.Context, commentID uuid.UUID) (comment *social_domain.Comment, err error)
	GetAllByPostID(ctx context.Context, postID uuid.UUID, limit, offset uint64) (comments []*social_domain.Comment, err error)
	// End Comment Table Queries //

}

type ITagsRepository interface {

	// Tag Table Commands //
	Create(ctx context.Context, tag *social_domain.Tag) (err error)
	DeleteByID(ctx context.Context, tagID uuid.UUID) (err error)
	// End Tag Table Commands //

	// Tag Table Queries //
	GetByID(ctx context.Context, tagID uuid.UUID) (tag *social_domain.Tag, err error)
	GetAll(ctx context.Context, limit, offset uint64) (tags []*social_domain.Tag, err error)
	// End Tag Table Queries //

}

type ISocialRepositories interface {
	PostsRepository() IPostsRepository
	CategoriesRepository() ICategoriesRepository
	CommentsRepository() ICommentsRepository
	TagsRepository() ITagsRepository
}
