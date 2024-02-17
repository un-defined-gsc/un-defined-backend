package v1

import (
	"github.com/ProjectMonWeb/API-Service/internal/config"
	"github.com/ProjectMonWeb/API-Service/internal/core"
	"github.com/ProjectMonWeb/API-Service/internal/delivery/http/error_handler"
	"github.com/ProjectMonWeb/API-Service/internal/delivery/http/response_types"
	"github.com/ProjectMonWeb/API-Service/internal/delivery/http/v1/private"
	"github.com/ProjectMonWeb/API-Service/internal/delivery/http/v1/public"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/swagger"
)

type V1Handler struct {
	adapter *core.CoreAdapter
}

func NewV1Handler(coreAdapter *core.CoreAdapter) *V1Handler {
	return &V1Handler{
		adapter: coreAdapter,
	}
}

func (h *V1Handler) Init(router fiber.Router, sessionStore *session.Store) {
	root := router.Group("/v1")
	root.Get("/", func(c *fiber.Ctx) error {
		return responseJson(200, response_types.HelloV1Route, nil)
	})
	// prodControl
	if !*config.GetConfig().App.Prod {
		root.Static("/docs/", "./docs")                    // default // air ile çalışıldığından ./docs/ olarak değiştirildi.
		root.Get("/swagger/*", swagger.New(swagger.Config{ // custom
			URL:          "/api/v1/docs/swagger.yaml",
			DocExpansion: "none",
		}))
	}
	// Init Fiber Session Store
	//---------------------------
	private.NewPrivateHandler(h.adapter, sessionStore, responseJson).Init(root)
	public.NewPublicHandler(h.adapter, sessionStore, responseJson).Init(root)
}

func responseJson(statusCode int, i18nStrKey string, data interface{}, dataCount ...uint64) error {
	var count uint64
	if len(dataCount) > 0 {
		count = dataCount[0]
	}
	return &error_handler.BaseResponse{
		StatusCode: statusCode,
		Message:    i18nStrKey,
		Data:       data,
		DataCount:  count,
	}
}
