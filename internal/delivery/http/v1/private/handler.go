package private

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/un-defined-gsc/un-defined-backend/internal/core"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

type PrivateHandler struct {
	coreAdapter   *core.CoreAdapter
	session_store *session.Store
	responseJson  func(statusCode int, i18nStrKey string, data interface{}, dataCount ...uint64) error
}

func NewPrivateHandler(
	coreAdapter *core.CoreAdapter,
	sessionStore *session.Store,
	responseJson func(statusCode int, i18nStrKey string, data interface{}, dataCount ...uint64) error,
) *PrivateHandler {
	return &PrivateHandler{
		coreAdapter:   coreAdapter,
		session_store: sessionStore,
		responseJson:  responseJson,
	}
}

func (h *PrivateHandler) Init(router fiber.Router) {
	root := router.Group("/private")

	root.Use(h.authMiddleware)

	root.Get("/", func(c *fiber.Ctx) error {
		return h.responseJson(200, "", nil)
	})
	h.initUserRoutes(root)
	h.initPostRoutes(root)
}

func (h *PrivateHandler) authMiddleware(c *fiber.Ctx) error {
	session, err := h.session_store.Get(c)
	if err != nil {
		return err
	}
	user := session.Get("user")
	if user == nil {
		return h.responseJson(401, response_types.ErrUnauthorized, nil)
	}
	userdto, ok := user.(domains.SessionDTO)
	if !ok {
		return h.responseJson(500, response_types.InternalServerError, "SessionDTO type assertion error")
	}
	if userdto.Disabled {
		return h.responseJson(403, response_types.ErrUnauthorized, nil)
	}
	if !userdto.EnabledSession {
		if !strings.Contains(c.Path(), "/private/user/enable/") {
			return h.responseJson(401, response_types.MFARequried, nil)
		}
	}
	if !userdto.EmailVerified {
		if !strings.Contains(c.Path(), "/private/user/me") {
			return h.responseJson(401, response_types.MailNotVerified, nil)
		}
	}
	c.Locals("user", userdto)
	return c.Next()
}
