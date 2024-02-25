package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
	social_repositories "github.com/un-defined-gsc/un-defined-backend/internal/repositories/postgresql_repositories/social"
)

type socialRepositories struct {
	postsRepository      social_ports.IPostsRepository
	categoriesRepository social_ports.ICategoriesRepository
	tagsRepository       social_ports.ITagsRepository
	imageRepository      social_ports.IImagesRepository
	likesRepository      social_ports.ILikesRepository
	commentsRepository   social_ports.ICommentsRepository
}

func NewSocialRepositories(dbpool *pgxpool.Pool) social_ports.ISocialRepositories {
	return &socialRepositories{
		postsRepository:      social_repositories.NewPostRepository(dbpool),
		categoriesRepository: social_repositories.NewCategoryRepository(dbpool),
		tagsRepository:       social_repositories.NewTagsRepository(dbpool),
		imageRepository:      social_repositories.NewImagesRepository(dbpool),
		likesRepository:      social_repositories.NewLikesRepository(dbpool),
	}
}

func (r *socialRepositories) PostsRepository() social_ports.IPostsRepository {
	return r.postsRepository
}

func (r *socialRepositories) CategoriesRepository() social_ports.ICategoriesRepository {
	return r.categoriesRepository
}

func (r *socialRepositories) TagsRepository() social_ports.ITagsRepository {
	return r.tagsRepository
}

func (r *socialRepositories) ImagesRepository() social_ports.IImagesRepository {
	return r.imageRepository
}

func (r *socialRepositories) LikesRepository() social_ports.ILikesRepository {
	return r.likesRepository
}

func (r *socialRepositories) CommentsRepository() social_ports.ICommentsRepository {
	return r.commentsRepository
}
