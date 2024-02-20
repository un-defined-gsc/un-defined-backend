package social_ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
)

type IPostsRepository interface {

	// Post Table Commands //
	Create(ctx context.Context, post *domains.CratePostDTO) (err error)
	Update(ctx context.Context, newPost *domains.UpdatePostDTO) (err error)
	DeleteByID(ctx context.Context, postID uuid.UUID, userID uuid.UUID) (err error)

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

type IImagesRepository interface {

	// Image Table Commands //
	Upload(ctx context.Context, image *social_domain.Image) (err error)
	DeleteByID(ctx context.Context, imageID uuid.UUID) (err error)
	// End Image Table Commands //

	// Image Table Queries //
	GetByID(ctx context.Context, imageID uuid.UUID) (image *social_domain.Image, err error)
	GetAll(ctx context.Context, limit, offset uint64) (images []*social_domain.Image, err error)
	// End Image Table Queries //

}

type ILikesRepository interface {

	// Like Table Commands //
	Like(ctx context.Context, like *social_domain.Like) (err error)
	UnLikeByID(ctx context.Context, likeID uuid.UUID) (err error)
	// End Like Table Commands //

}

type ISocialRepositories interface {
	PostsRepository() IPostsRepository
	CategoriesRepository() ICategoriesRepository
	CommentsRepository() ICommentsRepository
	TagsRepository() ITagsRepository
	ImagesRepository() IImagesRepository
	LikesRepository() ILikesRepository
}
