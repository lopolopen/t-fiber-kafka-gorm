//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log/slog"

	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/conf"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/gorm"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/gorm/repoimpl"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func wireApp(ctx context.Context, c *config.Config, orm conf.ORM, log *slog.Logger) (*fiber.App, error) {
	panic(wire.Build(
		repoimpl.ProviderSet,
		service.ProviderSet,
		gorm.NewGormDB,
		http.ProviderSet,
	))
}
