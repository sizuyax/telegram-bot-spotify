package telegram

import (

	tb "gopkg.in/tucnak/telebot.v2"
	
)

const (
	ModeHTML = tb.ModeHTML
)

var (
	menu     = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	selector = &tb.ReplyMarkup{}

	btnSearch       = menu.Text("🔍 Найти семью")
	btnBlock        = menu.Text("🚫 Заблокировать пользователя")
	btnUnblock      = menu.Text("🔓 Разблокировать пользователя")
	btnCheckBlocked = menu.Text("🔍 Проверить заблокированных пользователей")
	btnUnblockAll   = menu.Text("🔓 Разблокировать всех")
	btnReport       = menu.Text("🚨 Пожаловаться на пользователя")
)

func AllHandlers(b *tb.Bot) {

	menu.Reply(
		menu.Row(btnSearch, btnReport),
		menu.Row(btnBlock, btnUnblock),
		menu.Row(btnUnblockAll, btnCheckBlocked),
	)

	selector.Reply(
		selector.Row(btnSearch, btnReport),
		selector.Row(btnBlock, btnUnblock),
		selector.Row(btnUnblockAll, btnCheckBlocked),
	)

	ButtonHandlers(b)
	CommandHandlers(b)

}



