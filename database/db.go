package database

import (
	"log"

	"github.com/rajesh/personal-skill-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=12345 dbname=skill_tracker port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	db.AutoMigrate(&models.User{}, &models.Skill{}, &models.PracticeLog{})
}
