package http

import (
	"context"
	"log/slog"

	tfiberkafkagorm "github.com/lopolopen/t-fiber-kafka-gorm"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/dto"
	v1 "github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http/handlers/v1"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	"github.com/lopolopen/gap"
	"github.com/lopolopen/gap/broker/xkafka"
	"github.com/lopolopen/gap/storage/xgorm"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewPub, NewApp)

func NewPub(ctx context.Context, c *config.Config, db *gorm.DB, log *slog.Logger) gap.EventPublisher {
	if tfiberkafkagorm.HAVE_NOT_BEEN_DELETED_YET {
		return nil
	}

	pub := gap.NewEventPublisher(
		gap.WithDrain(ctx, 5),
		xgorm.UseGorm(
			xgorm.DB(db),
		),
		xkafka.UseKafka(
			xkafka.Brokers(c.Kafka.Brokers),
			xkafka.ConfigTopic(
			// xkafka.NumPartitions(4),
			// xkafka.ReplicationFactor(3),
			),
		),
		gap.UseDashboard(),
	)
	return pub
}

func NewApp(
	userSvc *service.UserSvc,
	pub gap.EventPublisher,
) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(dto.Err(err))
		},
	})
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	if pub != nil {
		app.All("/dashboard/*", adaptor.HTTPHandler(gap.NewDashboardHandler(pub)))
	}

	apiv1 := app.Group("/api/v1")

	users := apiv1.Group("/users")
	users.Get("", v1.Query(userSvc))

	if pub != nil {
		gap.Subscribe(
			gap.From(pub),
			gap.Inject(userSvc),
		)
	}
	return app
}
