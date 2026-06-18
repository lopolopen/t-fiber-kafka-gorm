package v1

import (
	"time"

	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/dto"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/timeout"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/query"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Query users.
// @Description	Query users by key.
// @Param			key	query	string	true	"Search key for User name"
// @Tags			User
// @Produce		json
// @Success		200	{object}	dto.Resp[[]result.User]
// @Failure		200	{object}	dto.Resp[any]
// @Router			/api/v1/users [get]
func QueryUsers(svc *service.UserSvc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Query("key")
		if key == "" {
			panic("key is empty")
		}
		users, err := svc.Query(c.UserContext(), query.UserQuery{Key: key})
		if err != nil {
			return err
		}
		return c.JSON(dto.OK(users))
	}
}

// @Summary		Query users. (Unsafe Timeout Demo)
// @Description	Query users by key.
// @Param			key	query	string	true	"Search key for User name"
// @Tags			User
// @Produce		json
// @Success		200	{object}	dto.Resp[[]result.User]
// @Failure		200	{object}	dto.Resp[any]
// @Router			/api/v1/users/unsafe-timeout [get]
func QueryUsersWithUnsafeTimeout(svc *service.UserSvc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		keepSafe := timeout.SecurityGuard(c)

		key := c.Query("key")
		if key == "" {
			panic("key is empty")
		}
		users, err := svc.Query(c.UserContext(), query.UserQuery{Key: key})
		time.Sleep(10 * time.Second)

		if err := keepSafe(); err != nil {
			return err
		}

		if err != nil {
			return err
		}
		return c.JSON(dto.OK(users))
	}
}
