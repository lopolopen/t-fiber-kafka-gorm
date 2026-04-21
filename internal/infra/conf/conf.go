package conf

import (
	"gorm.io/gorm/logger"
)

type MySQL struct {
	DSN          string `yaml:"dsn"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

type ORM struct {
	MySQL MySQL  `yaml:"mysql"`
	Level string `yaml:"logLevel"`
}

func (c ORM) GORMLogLevel() logger.LogLevel {
	switch c.Level {
	case "debug":
		return logger.Info
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	case "off":
		return logger.Silent
	}
	return logger.Info
}
