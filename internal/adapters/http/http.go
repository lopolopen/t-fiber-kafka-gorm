package http

import (
	"errors"
	"log/slog"

	"github.com/lopolopen/pkg/errorx"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/docs"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/dto"
	v1 "github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/handlers/v1"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"
	"github.com/lopolopen/t-fiber-kafka-gorm/pkg/schema"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	"github.com/lopolopen/gap"
)

var ProviderSet = wire.NewSet(NewApp)

func NewApp(
	c *config.Config,
	userSvc *service.UserSvc,
	pub gap.EventPublisher,
	logger *slog.Logger,
) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var e *errorx.Error
			if !errors.As(err, &e) {
				code := fiber.StatusInternalServerError
				reason := schema.ErrorReason_UNSPECIFIED.String()
				msg := err.Error()

				var fe *fiber.Error
				if errors.As(err, &fe) {
					code = fe.Code
					msg = fe.Message
				}
				e = errorx.New(code, reason, msg)
			}
			if e.Code >= 500 {
				logger.Error(err.Error())
			} else {
				logger.Debug(err.Error())
			}
			return c.Status(int(e.Code)).JSON(dto.Err(e))
		},
	})
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	docs.SwaggerInfo.Host = c.Swagger.Host
	docs.SwaggerInfo.BasePath = c.Swagger.BasePath
	docs.SwaggerInfo.Version += "-" + c.Env

	if pub != nil {
		app.All("/dashboard/*", adaptor.HTTPHandler(gap.NewDashboardHandler(pub)))
	}

	apiv1 := app.Group("/api/v1")

	users := apiv1.Group("/users")
	users.Get("", v1.QueryUsers(userSvc))

	if pub != nil {
		gap.Subscribe(
			gap.From(pub),
			gap.Inject(userSvc),
		)
	}
	return app
}
