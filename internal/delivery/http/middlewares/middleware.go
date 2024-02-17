package middlewares

import (
	"strings"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/un-defined-gsc/un-defined-backend/internal/config"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/error_handler"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
	"golang.org/x/text/language"
)

func InitMiddlewares(cfg *config.Config, store fiber.Storage) (mws []func(*fiber.Ctx) error) {
	// Fiber storage bağlantısı yapacağım unutma
	cors := cors.New(
		cors.Config{
			AllowOrigins:     strings.Join(cfg.Http.AllowedOrigins, ","),
			AllowMethods:     strings.Join(cfg.Http.AllowedMethods, ","),
			AllowHeaders:     strings.Join(cfg.Http.AllowedHeaders, ","),
			AllowCredentials: cfg.Http.AllowCredentials,
			ExposeHeaders:    strings.Join(cfg.Http.ExposedHeaders, ","),
			MaxAge:           cfg.Http.MaxAge,
		},
	)

	i18n := fiberi18n.New(&fiberi18n.Config{
		RootPath:        "./locales",
		DefaultLanguage: language.English,
		AcceptLanguages: []language.Tag{
			language.Turkish,
			language.English,
		},
		// LangHandler: , // session servisi yazılınca burası doldurulacak. şuan sadece accept-language header'ı ve lang query parametresi ile çalışıyor.
	})

	helmetMid := helmet.New(helmet.ConfigDefault) // helmet configleri yazılacak

	mws = append(mws, cors, i18n, helmetMid)

	if *cfg.App.Prod {
		limitter := limiter.New(limiter.Config{
			Max:     50,
			Storage: store,
			LimitReached: func(c *fiber.Ctx) error {
				return error_handler.ErrorHandler(c, &error_handler.BaseResponse{
					StatusCode: 429,
					Message:    response_types.ToManyRequest,
				})
			},
		})
		mws = append(mws, limitter)
	}
	return
}
