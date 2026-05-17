package http

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/lopolopen/pkg/errorx"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/dto"
	"github.com/lopolopen/t-fiber-kafka-gorm/pkg/schema/errx"
)

func HandlerError(logger *slog.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, theErr error) error {
		var err *errorx.Error
		var fErr *fiber.Error
		if errors.As(theErr, &fErr) {
			err = errx.FiberErr(fErr)
		} else if !errors.As(theErr, &err) {
			err = errx.ErrUnspecified.WithCause(theErr)
		}
		return c.Status(int(err.Code)).JSON(dto.Err(err))
	}
}
