package telegram

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"telegram-bot-spotify/backend"
	"telegram-bot-spotify/backend/database"
	"time"
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

	buttonHandlers(b)
	commandHandlers(b)

}

func buttonHandlers(b *tb.Bot) {
	b.Handle(&btnSearch, func(m *tb.Message) {
		id := m.Sender.ID
		username := m.Sender.Username

		users := database.GetUsers()

		ok := database.IsProfileExist(id)
		if ok {
			if len(users) <= 1 {
				b.Send(m.Sender, "Ты один в <b>базе данных</b> :(", ModeHTML)
				b.Send(m.Sender, "Попробуй еще раз через некоторое время нажав на команду /search", ModeHTML)
				return
			} else {
				profiles, err := database.GetProfile()
				if err != nil {
					b.Send(m.Sender, "Ошибка при обращении к серверу")
					return
				}

				replyMarkup := b.NewMarkup()

				for _, profile := range profiles {
					isReported := database.IsProfileBlocked(profile)

					if len(replyMarkup.InlineKeyboard) >= 5 {
						break
					} else if profile == username {
						continue
					} else if isReported {
						continue
					}

					url := fmt.Sprintf("https://t.me/%s", profile)
					btn := tb.InlineButton{Unique: profile, Text: profile, URL: url}

					replyMarkup.InlineKeyboard = append(replyMarkup.InlineKeyboard, []tb.InlineButton{btn})
				}

				find, _ := b.Send(m.Sender, "Мы найдем твою новую семью через: ", ModeHTML)

				three, _ := b.Send(m.Sender, "3", ModeHTML)

				time.Sleep(1 * time.Second)
				b.Delete(three)

				two, _ := b.Send(m.Sender, "2", ModeHTML)

				time.Sleep(1 * time.Second)
				b.Delete(two)

				one, _ := b.Send(m.Sender, "1", ModeHTML)

				time.Sleep(1 * time.Second)
				b.Delete(one)
				b.Delete(find)

				b.Send(m.Sender, "Вот твоя новая семья:", ModeHTML, &tb.SendOptions{ReplyMarkup: replyMarkup})
			}
		} else {
			b.Send(m.Sender, "Ты не занесен в базу данных", ModeHTML)
			b.Send(m.Sender, "Сейчас занесем тебя в базу данных!", ModeHTML)
			resp, err := backend.Client(id, username)
			if err != nil {
				b.Send(m.Sender, "Ошибка при обращении к серверу")
				return
			}
			b.Send(m.Sender, resp.Message)
			b.Send(m.Sender, "Теперь можешь нажать на команду /search чтобы найти свою <b>семью в Spotify</b>!", ModeHTML)
		}
	})

	b.Handle(&btnBlock, func(m *tb.Message) {
		users, _ := database.GetProfile()
		b.Send(m.Sender, "Напиши <b>username</b> пользователя которого хочешь заблокировать", ModeHTML)
		b.Handle(tb.OnText, func(m *tb.Message) {
			block := database.IsProfileBlocked(m.Text)
			if block {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> уже заблокирован", ModeHTML)
				return
			}

			found := false
			for _, user := range users {
				if m.Text == user {
					found = true
					err := database.BlockProfile(m.Sender.Username, m.Text)
					if err != nil {
						panic(err)
					}

					b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> заблокирован и больше не будет отображаться в поиске", ModeHTML)
					b.Send(m.Sender, "Чтобы разблокировать пользователя напиши /unblock", ModeHTML)
					break
				}
			}
			if !found {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> не найден", ModeHTML)
			}
		})
	})

	b.Handle(&btnUnblock, func(m *tb.Message) {
		users, _ := database.GetProfile()
		b.Send(m.Sender, "Напиши <b>username</b> пользователя которого хочешь разблокировать", ModeHTML)
		b.Handle(tb.OnText, func(m *tb.Message) {
			unblock := database.IsProfileBlocked(m.Text)
			if !unblock {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> не был заблокирован", ModeHTML)
				return
			}

			found := false
			for _, user := range users {
				if m.Text == user {
					found = true
					err := database.UnBlockProfile(m.Text)
					if err != nil {
						panic(err)
					}

					b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> разблокирован и будет отображаться в поиске", ModeHTML)
					break
				}
			}
			if !found {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> не найден", ModeHTML)
			}

		})
	})

	b.Handle(&btnCheckBlocked, func(m *tb.Message) {
		users, err := database.GetBlockedProfiles()
		if err != nil {
			b.Send(m.Sender, "Ошибка при обращении для получения списка заблокированных пользователей")
			return
		}
		if users == nil {
			b.Send(m.Sender, "Список заблокированных пользователей пуст")
			return
		}

		b.Send(m.Sender, "Список заблокированных пользователей:", ModeHTML)
		for _, user := range users {
			b.Send(m.Sender, "<b>"+user+"</b>", ModeHTML)
		}
	})

	b.Handle(&btnUnblockAll, func(m *tb.Message) {
		users, err := database.GetBlockedProfiles()
		if err != nil {
			b.Send(m.Sender, "Ошибка при обращении для получения списка заблокированных пользователей")
			return
		}
		if users == nil {
			b.Send(m.Sender, "Список заблокированных пользователей пуст")
			return
		}

		for _, user := range users {
			err = database.UnBlockProfile(user)
			if err != nil {
				return
			}
		}
		b.Send(m.Sender, "Все пользователи разблокированы", ModeHTML)
	})

	b.Handle(&btnReport, func(m *tb.Message) {
		users, _ := database.GetProfile()
		b.Send(m.Sender, "Напиши <b>username</b> пользователя на которого хочешь пожаловаться:", ModeHTML)
		b.Handle(tb.OnText, func(m *tb.Message) {
			found := false

			for _, user := range users {
				if m.Text == user {
					found = true
				}
			}
			if !found {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> не найден", ModeHTML)
				return
			}

			username := m.Text
			b.Send(m.Sender, "Теперь напиши причину жалобы:", ModeHTML)
			b.Handle(tb.OnText, func(m *tb.Message) {
				reportMessage := fmt.Sprintln(m.Text)
				b.Send(&tb.User{ID: 602974315}, "Пользователь <b>"+m.Sender.Username+"</b> пожаловался на пользователя <b>"+username+"</b> по причине: <b>"+reportMessage+"</b>", ModeHTML)
				b.Send(m.Sender, "Жалоба отправлена!", ModeHTML)
			})
		})
	})
}

func commandHandlers(b *tb.Bot) {
	var result *tb.Message
	b.Handle("/start", func(m *tb.Message) {
		reply := b.NewMarkup()
		reply.Inline(
			reply.Row(
				reply.Data("Продолжить", "button_clicked"),
			),
		)

		p := &tb.Photo{File: tb.FromDisk("./telegram/img/spotify.png")}
		b.Send(m.Sender, p)
		b.Send(m.Sender, "Тебя приветствует бот, который поможет найти людей для совместной подписки в <b>Spotify!</b>", &tb.SendOptions{ReplyMarkup: menu}, ModeHTML)
		result, _ = b.Send(m.Sender, "<i>Чтобы найти семью нажми продолжить!</i>", &tb.SendOptions{ReplyMarkup: reply}, ModeHTML)
	})

	b.Handle(&tb.InlineButton{Unique: "button_clicked"}, func(c *tb.Callback) {
		id := c.Sender.ID
		username := c.Sender.Username

		if username == "" {
			b.Send(c.Sender, "У тебя нет <b>username</b> в телеграмме, поэтому ты не можешь быть занесен в базу данных", ModeHTML)
			return
		}

		b.Delete(result)

		users := database.GetUsers()

		ok := database.IsProfileExist(id)
		if ok {
			b.Send(c.Sender, "Ты уже занесен в <b>базу данных</b> :)", ModeHTML)
			b.Send(c.Sender, "И поэтому тебе нужно нажать на команду /search чтобы найти свою <b>семью в Spotify</b>!", ModeHTML)
		} else {

			resp, err := backend.Client(id, username)
			if err != nil {
				b.Send(c.Sender, "Ошибка при обращении к серверу")
				return
			}

			if len(users) <= 1 {
				b.Send(c.Sender, "Ты один в <b>базе данных</b> :(", ModeHTML)
				b.Send(c.Sender, "Попробуй еще раз через некоторое время нажав на команду /search", ModeHTML)
				return
			} else {
				profiles, err := database.GetProfile()
				if err != nil {
					b.Send(c.Sender, "Ошибка при обращении к серверу")
					return
				}
				b.Send(c.Sender, resp.Message)
				wait, _ := b.Send(c.Sender, "Подожди немного мы почти нашли твою семью!", ModeHTML)

				time.Sleep(2 * time.Second)

				replyMarkup := b.NewMarkup()

				for _, profile := range profiles {
					if len(profiles) > 5 {
						break
					}
					if profile == username {
						continue
					}

					url := fmt.Sprintf("https://t.me/%s", profile)
					btn := tb.InlineButton{Unique: profile, Text: profile, URL: url}

					replyMarkup.InlineKeyboard = append(replyMarkup.InlineKeyboard, []tb.InlineButton{btn})
				}
				b.Delete(wait)

				b.Send(c.Sender, "Вот твоя новая семья:", ModeHTML, &tb.SendOptions{ReplyMarkup: replyMarkup})

				if err != nil {
					b.Send(c.Sender, "Ошибка при обращении к серверу")
					return
				}
			}
		}
	})

	b.Handle("/search", func(m *tb.Message) {
		id := m.Sender.ID
		username := m.Sender.Username

		users := database.GetUsers()

		ok := database.IsProfileExist(id)
		if ok {
			if len(users) <= 1 {
				b.Send(m.Sender, "Ты один в <b>базе данных</b> :(", ModeHTML)
				b.Send(m.Sender, "Попробуй еще раз через некоторое время нажав на команду /search", ModeHTML)
				return
			} else {
				profiles, err := database.GetProfile()
				if err != nil {
					b.Send(m.Sender, "Ошибка при обращении к серверу")
					return
				}

				replyMarkup := b.NewMarkup()

				for _, profile := range profiles {
					isReported := database.IsProfileBlocked(profile)

					if len(replyMarkup.InlineKeyboard) >= 5 {
						break
					} else if profile == username {
						continue
					} else if isReported {
						continue
					}

					url := fmt.Sprintf("https://t.me/%s", profile)
					btn := tb.InlineButton{Unique: profile, Text: profile, URL: url}

					replyMarkup.InlineKeyboard = append(replyMarkup.InlineKeyboard, []tb.InlineButton{btn})
				}

				find, _ := b.Send(m.Sender, "Мы найдем твою новую семью через: ", ModeHTML)

				three, _ := b.Send(m.Sender, "3", ModeHTML)

				time.Sleep(1 * time.Second)
				b.Delete(three)

				two, _ := b.Send(m.Sender, "2", ModeHTML)

				time.Sleep(1 * time.Second)
				b.Delete(two)

				one, _ := b.Send(m.Sender, "1", ModeHTML)

				time.Sleep(1 * time.Second)
				b.Delete(one)
				b.Delete(find)

				b.Send(m.Sender, "Вот твоя новая семья:", ModeHTML, &tb.SendOptions{ReplyMarkup: replyMarkup})
			}
		} else {
			b.Send(m.Sender, "Ты не занесен в базу данныхю", ModeHTML)
			b.Send(m.Sender, "Сейчас занесем тебя в базу данных!", ModeHTML)
			resp, err := backend.Client(id, username)
			if err != nil {
				b.Send(m.Sender, "Ошибка при обращении к серверу")
				return
			}
			b.Send(m.Sender, resp.Message)
			b.Send(m.Sender, "Теперь можешь нажать на команду /search чтобы найти свою <b>семью в Spotify</b>!", ModeHTML)
		}
	})

	b.Handle("/block", func(m *tb.Message) {
		users, _ := database.GetProfile()
		b.Send(m.Sender, "Напиши <b>username</b> пользователя которого хочешь заблокировать", ModeHTML)
		b.Handle(tb.OnText, func(m *tb.Message) {
			block := database.IsProfileBlocked(m.Text)
			if block {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> уже заблокирован", ModeHTML)
				return
			}

			found := false
			for _, user := range users {
				if m.Text == user {
					found = true
					err := database.BlockProfile(m.Sender.Username, m.Text)
					if err != nil {
						panic(err)
					}

					b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> заблокирован и больше не будет отображаться в поиске", ModeHTML)
					b.Send(m.Sender, "Чтобы разблокировать пользователя напиши /unblock", ModeHTML)
					break
				}
			}
			if !found {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> не найден", ModeHTML)
			}
		})
	})

	b.Handle("/unblock", func(m *tb.Message) {
		users, _ := database.GetProfile()
		b.Send(m.Sender, "Напиши <b>username</b> пользователя которого хочешь разблокировать", ModeHTML)
		b.Handle(tb.OnText, func(m *tb.Message) {
			unblock := database.IsProfileBlocked(m.Text)
			if !unblock {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> не был заблокирован", ModeHTML)
				return
			}

			found := false
			for _, user := range users {
				if m.Text == user {
					found = true
					err := database.UnBlockProfile(m.Text)
					if err != nil {
						panic(err)
					}

					b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> разблокирован и будет отображаться в поиске", ModeHTML)
					break
				}
			}
			if !found {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> не найден", ModeHTML)
			}
		})
	})

	b.Handle("/checkblocked", func(m *tb.Message) {
		users, err := database.GetBlockedProfiles()
		if err != nil {
			b.Send(m.Sender, "Ошибка при обращении для получения списка заблокированных пользователей")
			return
		}
		if users == nil {
			b.Send(m.Sender, "Список заблокированных пользователей пуст")
			return
		}

		b.Send(m.Sender, "Список заблокированных пользователей:", ModeHTML)
		for _, user := range users {
			b.Send(m.Sender, "<b>"+user+"</b>", ModeHTML)
		}
	})

	b.Handle("/unblockall", func(m *tb.Message) {
		users, err := database.GetBlockedProfiles()
		if err != nil {
			b.Send(m.Sender, "Ошибка при обращении для получения списка заблокированных пользователей")
			return
		}
		if users == nil {
			b.Send(m.Sender, "Список заблокированных пользователей пуст")
			return
		}

		for _, user := range users {
			err = database.UnBlockProfile(user)
			if err != nil {
				return
			}
		}
		b.Send(m.Sender, "Все пользователи разблокированы", ModeHTML)
	})

	b.Handle("/report", func(m *tb.Message) {
		users, _ := database.GetProfile()
		b.Send(m.Sender, "Напиши <b>username</b> пользователя на которого хочешь пожаловаться:", ModeHTML)
		b.Handle(tb.OnText, func(m *tb.Message) {
			found := false

			for _, user := range users {
				if m.Text == user {
					found = true
				}
			}
			if !found {
				b.Send(m.Sender, "Пользователь <b>"+m.Text+"</b> не найден", ModeHTML)
				return
			}

			username := m.Text
			b.Send(m.Sender, "Теперь напиши причину жалобы:", ModeHTML)
			b.Handle(tb.OnText, func(m *tb.Message) {
				reportMessage := fmt.Sprintln(m.Text)
				b.Send(&tb.User{ID: 602974315}, "Пользователь <b>"+m.Sender.Username+"</b> пожаловался на пользователя <b>"+username+"</b> по причине: <b>"+reportMessage+"</b>", ModeHTML)
				b.Send(m.Sender, "Жалоба отправлена!", ModeHTML)
			})
		})
	})
}
