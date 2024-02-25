package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PrivateHandler) initCommentRoutes(router fiber.Router) {
	comment := router.Group("/comment")
	comment.Post("/", h.CreateComment)
	comment.Delete("/:id<guid>", h.DeleteComment)

}

// @Tags Comment
// @Summary Create comment
// @Description Create comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param comment body domains.CommentDTO true "Comment"
// @Success 200 {object} error_handler.BaseResponse{data=domains.CommentDTO}
// @Router /private/comment [post]
func (h *PrivateHandler) CreateComment(c *fiber.Ctx) error {
	var comment domains.CommentDTO
	usersess := c.Locals("user").(domains.SessionDTO)
	if err := c.BodyParser(&comment); err != nil {
		return err
	}
	comment.UserID = *usersess.ID
	err := h.coreAdapter.SocialServices().CommentsService().CreateComment(c.Context(), &comment)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, comment)
}

// @Tags Comment
// @Summary Delete comment
// @Description Delete comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Comment ID"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/comment [delete]
func (h *PrivateHandler) DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	userID := c.Locals("user").(domains.SessionDTO).ID
	err = h.coreAdapter.SocialServices().CommentsService().DeleteComment(c.Context(), idUUID, *userID)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)

}
