package http

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/lopolopen/pkg/errorx"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/dto"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/timeout"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/pkg/x"
	"github.com/lopolopen/t-fiber-kafka-gorm/pkg/schema/errx"
)

func HandlerError(logger *slog.Logger) fiber.ErrorHandler {
	logger = x.SLogWithin(logger, HandlerError)

	return func(c *fiber.Ctx, theErr error) error {
		if errors.Is(theErr, timeout.ErrUnsafeFiberCtx) {
			logger.Debug("ctx is unsafe now")
			return nil
		}

		var err *errorx.Error
		var fErr *fiber.Error
		if errors.As(theErr, &fErr) {
			err = errx.FrameworkErr(fErr)
		} else if !errors.As(theErr, &err) {
			err = errx.ErrUnspecified.WithCause(theErr)
		}
		return c.Status(int(err.Code)).JSON(dto.Err(err))
	}
}
