package database

import (

	"github.com/jinzhu/gorm"
	"task5/models"
)

func AutoMigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{},&models.Photo{})
}