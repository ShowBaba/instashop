package db

import (
	"gorm.io/gorm"
	"instashop/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Product{},
		&models.OrderItem{},
	)
}
