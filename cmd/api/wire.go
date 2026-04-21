//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/http"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/applic/service"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/conf"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/gorm"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/gorm/repoimpl"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func wireApp(ctx context.Context, c *config.Config, orm conf.ORM) (*fiber.App, error) {
	panic(wire.Build(
		newLogger,
		repoimpl.ProviderSet,
		service.ProviderSet,
		gorm.NewGormDB,
		http.ProviderSet,
	))
}
