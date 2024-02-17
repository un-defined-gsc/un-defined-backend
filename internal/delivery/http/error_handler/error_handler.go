package error_handler

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
	"github.com/un-defined-gsc/un-defined-backend/pkg/validator_service"
)

type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	DataCount  uint64      `json:"data_count,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

func (r *BaseResponse) Error() string {
	return r.Message
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	base := &BaseResponse{}
	if errors.As(err, &base) {
		err.(*BaseResponse).Message = fiberi18n.MustLocalize(c, err.(*BaseResponse).Message) // Deneysel fiberi18n her an çalışmayabilir.
		return c.Status(err.(*BaseResponse).StatusCode).JSON(err)
	}
	if errors.As(err, &validator.ValidationErrors{}) {
		lang := c.AcceptsLanguages("tr", "en")
		errs := validator_service.ValidatorErrors(err, lang)
		return c.Status(400).JSON(
			&BaseResponse{
				StatusCode: 400,
				Message:    fiberi18n.MustLocalize(c, response_types.ValidationErrors),
				Errors:     errs,
			},
		)
	}
	fiberErr := &fiber.Error{}
	if errors.As(err, &fiberErr) {
		if fiberErr.Code == 404 {
			return c.Status(404).JSON(&BaseResponse{
				StatusCode: 404,
				Message:    fiberi18n.MustLocalize(c, response_types.FiberNotFound),
			})
		} else {
			return c.Status(err.(*fiber.Error).Code).JSON(&BaseResponse{
				StatusCode: err.(*fiber.Error).Code,
				Message:    err.(*fiber.Error).Message,
			})
		}
	}
	//database errors handle
	//

	serviceErr := &service_errors.ServiceError{}
	if errors.As(err, &serviceErr) {
		return c.Status(serviceErr.CodeSec).JSON(&BaseResponse{
			StatusCode: serviceErr.Code,
			Message:    fiberi18n.MustLocalize(c, serviceErr.Key),
		})
	}
	return c.Status(500).JSON(&BaseResponse{
		StatusCode: 500,
		Message:    "Internal Server Error (Unknown)", //ileride belki buna da key i18n keyi yazılabilir.
		Errors:     err.Error(),
	})
}
