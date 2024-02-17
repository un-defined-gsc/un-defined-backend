package public

import (
	"github.com/gofiber/fiber/v2"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PublicHandler) initCaptcaRoutes(router fiber.Router) {
	root := router.Group("/captcha")
	root.Get("/new", h.CaptchaNew)
	root.Get("/gen/:key", h.CaptchaGen)
}

// @Tags Captcha
// @Summary Captcha New
// @Description Captcha New
// @Accept json
// @Produce json
// @Success 200 {object} error_handler.BaseResponse{}
// @Router /public/captcha/new [get]
func (h *PublicHandler) CaptchaNew(c *fiber.Ctx) error {
	id := h.adapter.DepsServices().CaptchaService().New()
	return h.responseJson(200, response_types.RequestSuccess, fiber.Map{"captcha_id": id})
}

// @Tags Captcha
// @Summary Captcha Gen
// @Description Captcha Gen
// @Accept json
// @Produce json
// @Param key path string true "Key"
// @Success 200 {file} png
// @Router /public/captcha/gen/{key} [get]
func (h *PublicHandler) CaptchaGen(c *fiber.Ctx) error {
	key := c.Params("key")
	img, err := h.adapter.DepsServices().CaptchaService().GetImageBytes(key)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "image/png")
	return c.SendStream(&img)
}
