package db

import (
	"gorm.io/gorm"
	"ya41-56/internal/gophermart/models"
	"ya41-56/internal/shared/logger"
)

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		models.User{},
	); err != nil {
		logger.L().Info("auto migration to postgres failed")
		return err
	}

	return nil
}
