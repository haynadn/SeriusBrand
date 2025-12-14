package db

import (
	"seriusbrand/backend/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Order{},
		&models.UmkmPage{},
	)
}
