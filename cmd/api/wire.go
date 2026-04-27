//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log/slog"

	"github.com/lopolopen/gap/broker/xkafka"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/outbound"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/conf"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/gorm"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/gorm/repoimpl"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func wireApp(ctx context.Context, c *config.Config, k xkafka.Options, orm conf.ORM, log *slog.Logger) (*fiber.App, error) {
	panic(wire.Build(
		repoimpl.ProviderSet,
		service.ProviderSet,
		outbound.ProviderSet,
		gorm.NewGormDB,
		http.ProviderSet,
	))
}
