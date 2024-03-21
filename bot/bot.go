package bot

import (
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"telegram-bot-spotify/config"
	"telegram-bot-spotify/database"
	"telegram-bot-spotify/handlers"
)

func NewTelegramBot() error {
	logrus.Println("bot is running...")

	cfg := config.LoadEnv()
	pref := tb.Settings{
		Token:  cfg.TgBotApi,
		URL:    "https://api.telegram.org",
		Poller: &tb.LongPoller{Timeout: 10},
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		return err
	}

	database.InitDB(*cfg)
	handlers.AllHandlers(bot)

	bot.Start()
	return nil
}
