package http

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Timeout(h fiber.Handler, t time.Duration, errs ...error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		timeoutContext, cancel := context.WithTimeout(ctx.UserContext(), t)
		defer cancel()
		ctx.SetUserContext(timeoutContext)
		ch := make(chan error, 1)
		go func() {
			ch <- h(ctx)
		}()
		var err error
		select {
		case err = <-ch:
		case <-timeoutContext.Done():
			err = timeoutContext.Err()
		}
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				return fiber.ErrRequestTimeout
			}
			for i := range errs {
				if errors.Is(err, errs[i]) {
					return fiber.ErrRequestTimeout
				}
			}
			return err
		}
		return nil
	}
}
