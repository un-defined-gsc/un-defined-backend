package social_ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
)

type IPostsRepository interface {

	// Post Table Commands //
	Create(ctx context.Context, post *domains.CratePostDTO, categoryID uuid.UUID) (err error)
	Update(ctx context.Context, newPost *domains.UpdatePostDTO) (err error)
	DeleteByID(ctx context.Context, postID uuid.UUID, userID uuid.UUID) (err error)

	// End Post Table Commands //

	// Post Table Queries //
	GetByID(ctx context.Context, postID, userID uuid.UUID) (post domains.InPostDTO, err error)
	GetAll(ctx context.Context, limit, offset uint64) (posts []*domains.PostDTO, err error)
	GetByPostFilter(ctx context.Context, categoryID, userID uuid.UUID, tag string, limit, offset uint64) (posts []*domains.PostDTO, err error)

	// End Post Table Queries //

}

type ICategoriesRepository interface {

	// Category Table Queries //

	GetAll(ctx context.Context) (categories []*domains.CategoryDTO, err error)
	GetByName(ctx context.Context, name string) (category uuid.UUID, err error)
	// End Category Table Queries //

}

type ICommentsRepository interface {

	// Comment Table Commands //
	Create(ctx context.Context, comment *domains.CommentDTO) (err error)
	DeleteByID(ctx context.Context, commentID, userID uuid.UUID) (err error)
	// End Comment Table Commands //

	// Comment Table Queries //
	GetAllByPostID(ctx context.Context, postID uuid.UUID, limit, offset uint64) (comments []domains.ResCommentDTO, err error)
	// End Comment Table Queries //

}

type ITagsRepository interface {

	// Tag Table Commands //
	Create(ctx context.Context, tag *domains.CrateTagDTO) (err error)
	DeleteByID(ctx context.Context, tagID uuid.UUID) (err error)
	// End Tag Table Commands //

	// Tag Table Queries //
	GetByID(ctx context.Context, postID uuid.UUID) (tag []*domains.TagDTO, err error)
	GetAll(ctx context.Context, limit, offset uint64) (tags []*domains.TagDTO, err error)

	// End Tag Table Queries //

}

type IImagesRepository interface {

	// Image Table Commands //
	Create(ctx context.Context, image *social_domain.Image) (err error)
	DeleteByID(ctx context.Context, imageID uuid.UUID) (err error)
	Update(ctx context.Context, newImage *social_domain.Image) (err error)
	// End Image Table Commands //

	// Image Table Queries //
	GetByPath(ctx context.Context, imagePath string) (imageID uuid.UUID, err error)
	// End Image Table Queries //

}

type ILikesRepository interface {

	// Like Table Queries //
	GetLikesByPostID(ctx context.Context, like *domains.LikeDTO) (likes []*domains.LikeDTO, err error)
	// End Like Table Queries //
	// Like Table Commands //
	Like(ctx context.Context, like *domains.LikeDTO) (likeID uuid.UUID, err error)
	UnLike(ctx context.Context, like *domains.LikeDTO) (likeID uuid.UUID, err error)
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
