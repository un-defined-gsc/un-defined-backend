package social_service

import (
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type socialServices struct {
	postService       social_ports.IPostsService
	commentService    social_ports.ICommentsService
	imagesService     social_ports.IImagesService
	likesService      social_ports.ILikesService
	categoriesService social_ports.ICategoriesService
}

func NewSocialService(socialRepositories social_ports.ISocialRepositories, deps deps_ports.IDepsServices) social_ports.ISocialServices {
	return &socialServices{
		postService:       newPostService(socialRepositories, deps),
		commentService:    newCommentService(socialRepositories, deps),
		imagesService:     newImagesService(socialRepositories, deps),
		likesService:      newLikesService(socialRepositories, deps),
		categoriesService: newCategoriesService(socialRepositories, deps),
	}
}

func (s *socialServices) PostsService() social_ports.IPostsService {
	return s.postService
}
func (s *socialServices) CommentsService() social_ports.ICommentsService {
	return s.commentService
}
func (s *socialServices) ImagesService() social_ports.IImagesService {
	return s.imagesService
}
func (s *socialServices) LikesService() social_ports.ILikesService {
	return s.likesService
}
func (s *socialServices) CategoriesService() social_ports.ICategoriesService {
	return s.categoriesService
}
