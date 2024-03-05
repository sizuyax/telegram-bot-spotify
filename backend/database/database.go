package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"telegram-bot-spotify/backend/database/models"
)

var db *gorm.DB

func Initdb() {
	fmt.Println("database is running...")
	var err error
	dsn := "host=db_spotify user=admin password=pass dbname=db_spotify port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Profile{}, &models.ReportedProfile{})
	if err != nil {
		return
	}
}
