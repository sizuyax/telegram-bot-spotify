package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"telegram-bot-spotify/models"
)

func LoadEnv() *models.Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := models.Config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatal(err)
	}

	return &cfg
}
