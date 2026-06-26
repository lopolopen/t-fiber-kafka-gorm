package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lopolopen/t-fiber-kafka-gorm/cmd/api/config"
	_ "github.com/lopolopen/t-fiber-kafka-gorm/docs"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/pkg/confx"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/pkg/x"
	_ "go.uber.org/automaxprocs"
)

var commitSHA string
var f = flag.String("f", "etc/config.yaml", "config file")

// @title <app-name>-<org-name> API
// @version 1.0
// @description
// @contact.name Owner
// @contact.email user@example.com
// @host localhost:8080
// @BasePath /
func main() {
	flag.Parse()

	env := config.Env{
		Name:      os.Getenv("APP_ENV"),
		CommitSHA: commitSHA,
	}
	if env.Name == "" {
		env.Name = "prod"
	}

	var c config.Config
	confx.MustLoad(env.Name, *f, &c)

	if !env.IsProd() {
		cjson, _ := json.MarshalIndent(c, "", strings.Repeat(" ", 4))
		fmt.Printf("config of %s: %s\n", env.Name, string(cjson))
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := newLogger(c.Logger)

	var app *fiber.App
	app = x.Must(wireApp(ctx, &env, &c, c.Gap, c.Kafka, c.ORM, logger))

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
	slog.SetDefault(logger)
	return logger
}
