package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"telegram-bot-spotify/models"
)

var DB *gorm.DB

func InitDB(cfg models.Config) {
	logrus.Println("database is running...")

	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresHost,
		cfg.PostgresPorts,
		cfg.PostgresDB,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	if err = DB.AutoMigrate(&models.Profile{}, &models.ReportedProfile{}); err != nil {
		logrus.Fatal(err)
	}
}
