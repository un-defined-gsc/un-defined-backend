package public

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/un-defined-gsc/un-defined-backend/internal/core"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

type PublicHandler struct {
	adapter       *core.CoreAdapter
	session_store *session.Store
	responseJson  func(statusCode int, i18nStrKey string, data interface{}, dataCount ...uint64) error
}

func NewPublicHandler(
	adapter *core.CoreAdapter,
	sessionStore *session.Store,
	responseJson func(statusCode int, i18nStrKey string, data interface{}, dataCount ...uint64) error,
) *PublicHandler {
	return &PublicHandler{
		adapter:       adapter,
		session_store: sessionStore,
		responseJson:  responseJson,
	}
}

func (h *PublicHandler) Init(router fiber.Router) {
	root := router.Group("/public")

	root.Get("/", func(c *fiber.Ctx) error {
		return h.responseJson(200, response_types.HelloPublicRoute, nil)
	})
	h.initCaptcaRoutes(root)
	h.initUserRoutes(root)

}
