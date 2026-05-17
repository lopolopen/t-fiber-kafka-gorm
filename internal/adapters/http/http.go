package http

import (
	"log/slog"
	"time"

	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/docs"
	v1 "github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/handlers/v1"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"

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
		ErrorHandler: HandlerError(logger),
		BodyLimit:    4 * 1024 * 1024,
	})
	if c.IsProd() && c.Timeout > 0 {
		app.Use(Timeout(func(c *fiber.Ctx) error {
			return c.Next()
		}, time.Duration(c.Timeout)*time.Millisecond))
	}
	app.Use(recover.New(recover.Config{
		EnableStackTrace: !c.IsProd(),
	}))
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
