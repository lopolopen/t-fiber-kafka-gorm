package timeout

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

var ErrUnsafeFiberCtx = errors.New("unsafe fiber context")

const unsafeCtxKey = "unsafeCtx"

func SetUnsafeCtx(c *fiber.Ctx, ctx context.Context) {
	c.Locals(unsafeCtxKey, ctx)
}

func SecurityGuard(c *fiber.Ctx) func() error {
	ctx, ok := c.Locals(unsafeCtxKey).(context.Context)

	return func() error {
		if !ok {
			return nil
		}
		select {
		case <-ctx.Done():
			return ErrUnsafeFiberCtx
		default:
		}
		return nil
	}
}

// UnsafeTimeout wraps a fiber.Handler with a maximum execution time limit.
// ⚠️ WARNING: This middleware is CONCURRENCY UNSAFE and should be used with extreme caution.
// Dangerous reasons:
//  1. Data Race: It runs the user handler 'h' in a separate background goroutine. Since
//     fiber.Ctx is NOT thread-safe, any concurrent access or delayed modification to the ctx
//     (or its underlying response buffers) will cause data races and unexpected panics.
//  2. Context/Pool Pollution: When the timeout 't' expires, this middleware returns a 408
//     Request Timeout immediately to the client, allowing Fiber to recycle the fiber.Ctx back
//     into the sync.Pool. However, the background goroutine executing 'h' may STILL be running.
//     If that ghost goroutine later mutates or writes to the recycled fiber.Ctx, it will
//     silently pollute or corrupt the context/data of another user's incoming request.
func UnsafeTimeout(h fiber.Handler, t time.Duration, errs ...error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		timeoutContext, cancel := context.WithTimeout(ctx.UserContext(), t)
		defer cancel()
		ctx.SetUserContext(timeoutContext)
		SetUnsafeCtx(ctx, timeoutContext)
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
