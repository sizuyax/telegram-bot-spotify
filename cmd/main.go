package main

import (
	"github.com/sirupsen/logrus"
	"telegram-bot-spotify/bot"
)

func main() {
	if err := bot.NewTelegramBot(); err != nil {
		logrus.Fatal(err)
	}
}
