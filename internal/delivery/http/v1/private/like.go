package private

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PrivateHandler) initLikeRoutes(router fiber.Router) {
	like := router.Group("/like")
	like.Post("/:id<guid>", h.Like)
	like.Delete("/:id<guid>", h.UnLike)

}

// @Tags Like
// @Summary Like post
// @Description Like post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Post ID"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/like/{id} [post]
func (h *PrivateHandler) Like(c *fiber.Ctx) error {
	postID := c.Params("id")
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return err
	}
	userID := c.Locals("user").(domains.SessionDTO).ID
	fmt.Println(userID)
	err = h.coreAdapter.SocialServices().LikesService().Like(c.Context(), &domains.LikeDTO{
		UserID: *userID,
		PostID: postUUID,
	})
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Like
// @Summary UnLike post
// @Description UnLike post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Post ID"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/like/{id} [delete]
func (h *PrivateHandler) UnLike(c *fiber.Ctx) error {
	postID := c.Params("id")
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return err
	}
	userID := c.Locals("user").(domains.SessionDTO).ID

	err = h.coreAdapter.SocialServices().LikesService().UnLike(c.Context(), &domains.LikeDTO{
		UserID: *userID,
		PostID: postUUID,
	})
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}
