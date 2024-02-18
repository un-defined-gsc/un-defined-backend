package server

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/un-defined-gsc/un-defined-backend/internal/config"
)

type Server struct {
	app  *fiber.App
	port string
}

func NewServer(cfg *config.Config, errHandler ...func(c *fiber.Ctx, err error) error) *Server {
	readTimeout, err := time.ParseDuration(cfg.Http.ReadTimeout)
	if err != nil {
		readTimeout = 15 * time.Second
	}
	writeTimeout, err := time.ParseDuration(cfg.Http.WriteTimeout)
	if err != nil {
		writeTimeout = 15 * time.Second
	}
	fiberConfig := fiber.Config{
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		ServerHeader: cfg.App.Name + "/" + cfg.App.Version,
		ProxyHeader:  cfg.Http.ProxyHeader,
	}
	if len(errHandler) > 0 {
		fiberConfig.ErrorHandler = errHandler[0]
	}
	return &Server{
		app:  fiber.New(fiberConfig),
		port: cfg.Http.Port,
	}
}

func (s *Server) Run(apiApp *fiber.App) error {
	s.app.Mount("/", apiApp)
	return s.app.Listen(":" + s.port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx2, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	return s.app.ShutdownWithContext(ctx2)
}
