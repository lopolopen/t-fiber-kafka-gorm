package v1

import (
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/dto"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/query"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Query users.
// @Description	Query users by key.
// @Param			key	query	string	true	"Search key for User name"
// @Tags			User
// @Produce		json
// @Success		200	{object}	dto.Resp
// @Failure		200	{object}	dto.Resp{data=nil}
// @Router			/api/v1/users [get]
func QueryUsers(svc *service.UserSvc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Query("key")
		if key == "" {
			panic("key is empty")
		}
		users, err := svc.Query(c.Context(), query.UserQuery{Key: key})
		if err != nil {
			return err
		}
		return c.JSON(dto.OK(users))
	}
}
