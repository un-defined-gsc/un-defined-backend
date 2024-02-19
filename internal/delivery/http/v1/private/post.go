package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PrivateHandler) initPostRoutes(root fiber.Router) {
	post := root.Group("/post")
	post.Post("/", h.CreatePost)
	post.Put("/", h.UpdatePost)
	post.Delete("/:id<guid>", h.DeletePost)
	post.Get("/:id<guid>", h.GetPost)
	post.Get("/", h.GetPosts)
}

// @Tags Post
// @Summary Create post
// @Description Create post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param post body social_domain.Post true "Post"
// @Success 200 {object} error_handler.BaseResponse{data=social_domain.Post}
// @Router /private/post [post]
func (h *PrivateHandler) CreatePost(c *fiber.Ctx) error {
	var post social_domain.Post
	if err := c.BodyParser(&post); err != nil {
		return err
	}
	err := h.coreAdapter.SocialServices().PostsService().CreatePost(c.Context(), &post)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, post)
}

// @Tags Post
// @Summary Update post
// @Description Update post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param post body social_domain.Post true "Post"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/post [put]
func (h *PrivateHandler) UpdatePost(c *fiber.Ctx) error {
	var post social_domain.Post
	if err := c.BodyParser(&post); err != nil {
		return err
	}
	err := h.coreAdapter.SocialServices().PostsService().UpdatePost(c.Context(), &post)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Post
// @Summary Delete post
// @Description Delete post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Post ID"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/post/{id} [delete]
func (h *PrivateHandler) DeletePost(c *fiber.Ctx) error {
	postID := c.Params("id")
	newPostID := uuid.MustParse(postID)
	err := h.coreAdapter.SocialServices().PostsService().DeletePost(c.Context(), newPostID)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Post
// @Summary Get post
// @Description Get post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Post ID"
// @Success 200 {object} error_handler.BaseResponse{data=social_domain.Post}
// @Router /private/post/{id} [get]
func (h *PrivateHandler) GetPost(c *fiber.Ctx) error {
	postID := c.Params("id")
	newPostID := uuid.MustParse(postID)
	post, err := h.coreAdapter.SocialServices().PostsService().GetPost(c.Context(), newPostID)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, post)
}

// @Tags Post
// @Summary Get posts
// @Description Get posts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} error_handler.BaseResponse{data=social_domain.Post}
// @Router /private/post [get]
func (h *PrivateHandler) GetPosts(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	offset := c.QueryInt("offset")

	posts, err := h.coreAdapter.SocialServices().PostsService().GetPosts(c.Context(), uint64(limit), uint64(offset))
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, posts)
}
