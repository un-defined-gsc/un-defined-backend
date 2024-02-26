package private

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PrivateHandler) initPostRoutes(root fiber.Router) {
	post := root.Group("/post")
	post.Post("/", h.CreatePost)
	post.Put("/", h.UpdatePost)
	post.Delete("/:id<guid>", h.DeletePost)
	post.Get("/", h.GetPosts)
	post.Get("/:id<guid>", h.GetPost)
	post.Get("/filter", h.GetPostsByFilter)

}

// @Tags Post
// @Summary Create post
// @Description Create post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param post body domains.CratePostDTO true "Post"
// @Success 200 {object} error_handler.BaseResponse{data=domains.CratePostDTO}
// @Router /private/post [post]
func (h *PrivateHandler) CreatePost(c *fiber.Ctx) error {
	var post domains.CratePostDTO
	usersess := c.Locals("user").(domains.SessionDTO)
	if err := c.BodyParser(&post); err != nil {
		return err
	}
	post.UserID = *usersess.ID
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
// @Param post body domains.UpdatePostDTO true "Post"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/post [put]
func (h *PrivateHandler) UpdatePost(c *fiber.Ctx) error {
	var post domains.UpdatePostDTO
	if err := c.BodyParser(&post); err != nil {
		return err
	}
	userID := c.Locals("user").(domains.SessionDTO).ID
	post.UserID = *userID
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
	newPostID, err := uuid.Parse(postID)
	if err != nil {
		return err
	}
	userID := c.Locals("user").(domains.SessionDTO).ID
	err = h.coreAdapter.SocialServices().PostsService().DeletePost(c.Context(), newPostID, *userID)
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
// @Success 200 {object} error_handler.BaseResponse{data=domains.InPostDTO}
// @Router /private/post/{id} [get]
func (h *PrivateHandler) GetPost(c *fiber.Ctx) error {
	postID := c.Params("id")
	limit := c.QueryInt("limit")
	if limit == 0 {
		limit = 10
	}
	offset := c.QueryInt("offset")
	newPostID, err := uuid.Parse(postID)
	if err != nil {
		return err
	}
	userID := c.Locals("user").(domains.SessionDTO).ID
	post, err := h.coreAdapter.SocialServices().PostsService().GetPost(c.Context(), newPostID, *userID, uint64(limit), uint64(offset))
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
// @Success 200 {object} error_handler.BaseResponse{data=domains.PostDTO}
// @Router /private/post [get]
func (h *PrivateHandler) GetPosts(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	if limit == 0 {
		limit = 10
	}
	offset := c.QueryInt("offset")

	posts, err := h.coreAdapter.SocialServices().PostsService().GetPosts(c.Context(), uint64(limit), uint64(offset))
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, posts)
}

// @Tags Post
// @Summary Get posts by filter
// @Description Get posts by filter
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param categoryID query string false "Category ID"
// @Param userID query string false "User ID"
// @Param tag query string false "Tag"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} error_handler.BaseResponse{data=domains.PostDTO}
// @Router /private/post/filter [get]
func (h *PrivateHandler) GetPostsByFilter(c *fiber.Ctx) error {
	categoryID := c.Query("categoryID")
	userID := c.Query("userID")
	tag := c.Query("tag")
	limit := c.QueryInt("limit")
	if limit == 0 {
		limit = 10
	}
	offset := c.QueryInt("offset")
	newCategoryID, _ := uuid.Parse(categoryID)

	newUserID, _ := uuid.Parse(userID)
	fmt.Println(newCategoryID, newUserID, tag, limit, offset)

	posts, err := h.coreAdapter.SocialServices().PostsService().GetPostByFilter(c.Context(), newCategoryID, newUserID, tag, uint64(limit), uint64(offset))
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, posts)
}
