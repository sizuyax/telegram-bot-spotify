package main

import (
	"telegram-bot-spotify/telegram"
)

func main() {
	err := telegram.NewTelegramBot()
	if err != nil {
		panic(err)
	}
}
