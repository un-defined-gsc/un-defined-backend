package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PrivateHandler) initRoadmapRoutes(root fiber.Router) {
	roadmap := root.Group("/roadmap")
	category := roadmap.Group("/category")
	category.Get("/", h.RoadmapGetCategories)
	category.Post("/", h.RoadmapCreateCategory)
	category.Put("/:id<guid>", h.RoadmapUpdateCategory)
	category.Delete("/:id<guid>", h.RoadmapDeleteCategory)

	maps := roadmap.Group("/maps")
	maps.Get("/:categor_id<guid>?", h.GetRoadmaps)

}

// @Tags Roadmap/Category
// @Summary Get categories
// @Description Get categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param name query string false "Name"
// @Param id query string false "ID"
// @Success 200 {object} error_handler.BaseResponse{data=[]roadmap_domain.Category}
// @Router /private/roadmap/categories [get]
func (h *PrivateHandler) RoadmapGetCategories(c *fiber.Ctx) error {
	id := c.Query("id")
	name := c.Query("name")
	categoryID := uuid.Nil //Burada bir patlama olcak gibi
	if id != "" {
		categoryID, _ = uuid.Parse(id)
	}
	categories, err := h.coreAdapter.RoadmapServices().CategoryService().SearchCategory(c.Context(), name, categoryID)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, categories)
}

// @Tags Roadmap/Category
// @Summary Create category
// @Description Create category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category body roadmap_domain.Category true "Category"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/roadmap/categories [post]
func (h *PrivateHandler) RoadmapCreateCategory(c *fiber.Ctx) error {
	var category roadmap_domain.Category
	if err := c.BodyParser(&category); err != nil {
		return err
	}
	err := h.coreAdapter.RoadmapServices().CategoryService().CreateCategory(c.Context(), &category)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Roadmap/Category
// @Summary Delete category
// @Description Delete category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "ID"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/roadmap/categories/{id} [delete]
func (h *PrivateHandler) RoadmapDeleteCategory(c *fiber.Ctx) error {
	categoryID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}
	err = h.coreAdapter.RoadmapServices().CategoryService().DeleteCategory(c.Context(), categoryID)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Roadmap/Category
// @Summary Update category
// @Description Update category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "ID"
// @Param category body roadmap_domain.Category true "Category"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/roadmap/categories/{id} [put]

func (h *PrivateHandler) RoadmapUpdateCategory(c *fiber.Ctx) error {
	var category roadmap_domain.Category
	if err := c.BodyParser(&category); err != nil {
		return err
	}
	category.ID, _ = uuid.Parse(c.Params("id"))
	err := h.coreAdapter.RoadmapServices().CategoryService().UpdateCategory(c.Context(), &category)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

//

func (h *PrivateHandler) GetRoadmaps(c *fiber.Ctx) error {
	category_id := uuid.Nil
	if c.Params("categor_id") != "" {
		category_id = uuid.MustParse(c.Params("category_id"))
	}
	roadmaps, err := h.coreAdapter.RoadmapServices().RoadmapService().SearchRoadmap(c.Context(), category_id)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, roadmaps)
}
