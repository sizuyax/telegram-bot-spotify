package telegram

import (
	"fmt"
	"telegram-bot-spotify/backend/database"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func ButtonHandlers(b *tb.Bot) {
	b.Handle(&btnSearch, func(m *tb.Message) {
		id := m.Sender.ID
		username := m.Sender.Username
	search:
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
			notInBD, _ := b.Send(m.Sender, "Ты не занесен в базу данных", ModeHTML)
			time.Sleep(2 * time.Second)
			b.Delete(notInBD)
			process, _ := b.Send(m.Sender, "Сейчас занесем тебя в базу данных!", ModeHTML)
			err := database.AddProfileToDB(id, username)
			if err != nil {
				b.Send(m.Sender, "Ошибка при обращении к серверу")
				return
			}
			time.Sleep(2 * time.Second)
			b.Delete(process)
			b.Send(m.Sender, "Тебя занесли в базу данных.")

			goto search
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
