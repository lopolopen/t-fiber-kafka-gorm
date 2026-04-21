package config

import (
	"log/slog"

	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/conf"

	"github.com/lopolopen/gap/broker/xkafka"
)

type Config struct {
	Env    string         `yaml:"env"`
	Port   int            `yaml:"port"`
	Bind   string         `yaml:"bind"`
	Logger Logger         `yaml:"logger"`
	ORM    conf.ORM       `yaml:"orm"`
	Kafka  xkafka.Options `yaml:"kafka"`
}

type Logger struct {
	Level string `yaml:"level"`
	JSON  bool   `yaml:"json"`
}

func (l Logger) LogLevel() slog.Level {
	switch l.Level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "off":
		return slog.LevelError + 1
	}
	return slog.LevelInfo
}
