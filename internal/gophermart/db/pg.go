package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ya41-56/internal/shared/logger"
)

type InitPostgresConfig struct {
	DSN             string
	IsFireMigration bool
}

func InitPostgres(config *InitPostgresConfig) *gorm.DB {
	if config.DSN == "" {
		return nil
	}

	gormDB, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		logger.L().Info("gorm connection to postgres failed")
		return nil
	}
	if config.IsFireMigration {
		migrate(gormDB)
	}

	logger.L().Info("connected to postgres")
	return gormDB
}
