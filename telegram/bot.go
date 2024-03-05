package telegram

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"telegram-bot-spotify/backend/database"
	"telegram-bot-spotify/utils"
)

func NewTelegramBot() error {
	fmt.Println("bot is running...")

	err := utils.LoadEnv()
	if err != nil {
		return err
	}

	tokenBot := os.Getenv("TELEGRAM_TOKEN_BOT")

	pref := tb.Settings{
		Token:  tokenBot,
		URL:    "https://api.telegram.org",
		Poller: &tb.LongPoller{Timeout: 10},
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		return err
	}

	AllHandlers(b)
	go database.Initdb()

	b.Start()
	return nil
}
