package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	_ "github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/docs"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/pkg/x"
	_ "go.uber.org/automaxprocs"
)

var f = flag.String("f", "etc/config.yaml", "config file")
var k = koanf.New(".")

// @title Example API
// @version 1.0
// @description This is the Example API documentation.
// @contact.name UserName
// @contact.email user@example.com
// @host localhost:8080
// @BasePath /
func main() {
	flag.Parse()

	if err := k.Load(file.Provider(*f), yaml.Parser()); err != nil {
		panic(err)
	}

	var c config.Config
	if err := k.Unmarshal("", &c); err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := newLogger(c.Logger)

	var app *fiber.App
	app = x.Must(wireApp(ctx, &c, c.ORM, logger))

	go func() {
		err := app.Listen(fmt.Sprintf("%s:%d", c.Bind, c.Port))
		if err != nil {
			logger.Error("server stop listening with error", slog.Any("err", err))
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down gracefully, press Ctrl+C again to force")

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		logger.Warn("forced shutdown by user, exiting immediately")
		os.Exit(1)
	}()

	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		logger.Warn("server shutdown with error", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("server shutdown successfully")
}

func newLogger(c config.Logger) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: c.LogLevel(),
	}
	var h slog.Handler
	if c.JSON {
		h = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		h = slog.NewTextHandler(os.Stdout, opts)
	}
	logger := slog.New(h)
	return logger
}
