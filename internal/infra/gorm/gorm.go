package gorm

import (
	"fmt"

	tfiberkafkagorm "github.com/lopolopen/t-fiber-kafka-gorm"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/conf"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/infra/gorm/migrate"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// NewGormDB .
func NewGormDB(c conf.ORM) *gorm.DB {
	if tfiberkafkagorm.HAVE_NOT_BEEN_DELETED_YET {
		return nil
	}

	db, err := gorm.Open(mysql.Open(c.MySQL.DSN), &gorm.Config{
		Logger:                                   logger.Default.LogMode(c.GORMLogLevel()),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		DisableForeignKeyConstraintWhenMigrating: true,
		TranslateError:                           true,
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %s", err))

	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(c.MySQL.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MySQL.MaxOpenConns)

	err = migrate.RunMigrate(db)
	if err != nil {
		panic(fmt.Errorf("failed to migrate: %s", err))
	}
	return db
}
