package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PrivateHandler) initCategoryRoutes(router fiber.Router) {
	category := router.Group("/category")
	category.Get("/", h.GetCategories)

}

// @Tags Category
// @Summary Get categories
// @Description Get categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} error_handler.BaseResponse{data=[]domains.CategoryDTO}
// @Router /private/category [get]
func (h *PrivateHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.coreAdapter.SocialServices().CategoriesService().GetCategories(c.Context())
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, categories)
}
