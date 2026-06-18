package http

import (
	"log/slog"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/docs"
	v1 "github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/handlers/v1"
	tout "github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/timeout"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"

	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	"github.com/lopolopen/gap"
)

var ProviderSet = wire.NewSet(NewApp)

func NewApp(
	e *config.Env,
	c *config.Config,
	userSvc *service.UserSvc,
	pub gap.EventPublisher,
	logger *slog.Logger,
) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: HandlerError(logger),
		BodyLimit:    4 * 1024 * 1024,
	})
	if !e.IsProd() && c.Timeout > 0 {
		app.Use(timeout.NewWithContext(func(c *fiber.Ctx) error {
			return c.Next()
		}, time.Duration(c.Timeout)*time.Millisecond))
	}
	app.Use(recover.New(recover.Config{
		EnableStackTrace: !e.IsProd(),
	}))
	corsConf := cors.ConfigDefault
	if len(c.CORS.AllowOrigins) > 0 {
		corsConf.AllowOrigins = strings.Join(c.CORS.AllowOrigins, ",")
	}
	if len(c.CORS.AllowHeaders) > 0 {
		corsConf.AllowHeaders = strings.Join(c.CORS.AllowHeaders, ",")
	}
	app.Use(cors.New(corsConf))

	app.Get("/swagger/*", swagger.HandlerDefault)
	docs.SwaggerInfo.Host = c.Swagger.Host
	docs.SwaggerInfo.BasePath = c.Swagger.BasePath
	docs.SwaggerInfo.Title += "-" + e.Name
	if e.CommitSHA != "" {
		docs.SwaggerInfo.Version += "-" + e.CommitSHA
	}

	if pub != nil {
		app.All("/dashboard/*", adaptor.HTTPHandler(gap.NewDashboardHandler(pub)))
	}

	apiv1 := app.Group("/api/v1")

	users := apiv1.Group("/users")
	users.Get("", v1.QueryUsers(userSvc))
	users.Get("unsafe-timeout", tout.UnsafeTimeout(v1.QueryUsersWithUnsafeTimeout(userSvc), 1*time.Second))

	if pub != nil {
		gap.Subscribe(
			gap.From(pub),
			gap.Inject(userSvc),
		)
	}
	return app
}
