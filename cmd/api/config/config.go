package config

import (
	"log/slog"

	"github.com/lopolopen/gap/broker/xkafka"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/conf"
)

type Env struct {
	Name      string
	CommitSHA string
	Version   string
}

func (e *Env) IsProd() bool {
	return e.Name == "prod"
}

type Config struct {
	Port    int            `yaml:"port"`
	Bind    string         `yaml:"bind"`
	Timeout int64          `yaml:"timeout"`
	CORS    CORS           `yaml:"cors"`
	Swagger Swagger        `yaml:"swagger"`
	Gap     Gap            `yaml:"gap"`
	Logger  Logger         `yaml:"logger"`
	ORM     conf.ORM       `yaml:"orm"`
	Kafka   xkafka.Options `yaml:"kafka"`
}

type CORS struct {
	AllowOrigins []string `yaml:"allow_origins"`
	AllowHeaders []string `yaml:"allow_headers"`
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

type Swagger struct {
	Host     string `yaml:"host"`
	BasePath string `yaml:"base_path"`
}

type Gap struct {
	Location string `yaml:"location"`
}
