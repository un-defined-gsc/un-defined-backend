package http

import (
	"github.com/ProjectMonWeb/API-Service/internal/config"
	"github.com/ProjectMonWeb/API-Service/internal/core"
	"github.com/ProjectMonWeb/API-Service/internal/delivery/http/session"
	v1 "github.com/ProjectMonWeb/API-Service/internal/delivery/http/v1"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	adapter *core.CoreAdapter
}

func NewHandler(adapter *core.CoreAdapter) *Handler {
	return &Handler{
		adapter: adapter,
	}
}

func (h *Handler) Init(cfg *config.Config, fiberStore fiber.Storage, middlewares ...func(*fiber.Ctx) error) *fiber.App {
	app := fiber.New()
	//init middlewares

	for i := range middlewares {
		app.Use(middlewares[i])
	}

	root := app.Group("/api")

	// init routes
	sessionStore := session.NewSessionStore(fiberStore)

	v1.NewV1Handler(h.adapter).Init(root, sessionStore)

	return app
}
