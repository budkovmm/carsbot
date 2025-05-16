package bot

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v4"
)

func RegisterHandlers(b *telebot.Bot) {
	b.Handle("/start", func(c telebot.Context) error {
		msg := "👋 Здравствуйте! Я помогу вам оформить договор купли-продажи автомобиля.\n\n" +
			"Чтобы начать, просто напишите любое сообщение или нажмите кнопку ниже."
		markup := &telebot.ReplyMarkup{}
		btnStart := markup.Text("Оформить договор")
		markup.Reply(markup.Row(btnStart))
		return c.Send(msg, markup)
	})

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		userID := c.Sender().ID
		state, ok := userStates[userID]

		if !ok {
			if c.Text() == "Оформить договор" {
				state = &UserState{Step: 0}
				userStates[userID] = state
			} else {
				return nil // Игнорируем прочие сообщения
			}
		}

		switch state.Step {
		case 0:
			state.Step = 1
			return c.Send("Здравствуйте! Давайте оформим договор купли-продажи авто.\nВведите ФИО продавца:")
		case 1:
			state.SellerName = c.Text()
			state.Step = 2
			return c.Send("Введите ФИО покупателя:")
		case 2:
			state.BuyerName = c.Text()
			state.Step = 3
			return c.Send("Введите VIN автомобиля:")
		case 3:
			state.VIN = c.Text()
			state.Step = 4
			return c.Send("Введите марку и модель автомобиля:")
		case 4:
			state.BrandModel = c.Text()
			state.Step = 5
			return c.Send("Введите год выпуска автомобиля:")
		case 5:
			state.Year = c.Text()
			state.Step = 6
			return c.Send("Введите цвет автомобиля:")
		case 6:
			state.Color = c.Text()
			state.Step = 7
			return c.Send("Введите стоимость автомобиля (в рублях):")
		case 7:
			state.Price = c.Text()
			state.Step = 8
			return c.Send("Введите дату сделки (например, 12.06.2024):")
		case 8:
			state.Date = c.Text()
			state.Step = 9
			return c.Send("Введите город сделки:")
		case 9:
			state.City = c.Text()
			// Выводим все данные списком
			var sb strings.Builder
			sb.WriteString("Спасибо! Вот все данные, которые вы ввели:\n")
			sb.WriteString(fmt.Sprintf("ФИО продавца: %s\n", state.SellerName))
			sb.WriteString(fmt.Sprintf("ФИО покупателя: %s\n", state.BuyerName))
			sb.WriteString(fmt.Sprintf("VIN: %s\n", state.VIN))
			sb.WriteString(fmt.Sprintf("Марка/модель: %s\n", state.BrandModel))
			sb.WriteString(fmt.Sprintf("Год выпуска: %s\n", state.Year))
			sb.WriteString(fmt.Sprintf("Цвет: %s\n", state.Color))
			sb.WriteString(fmt.Sprintf("Стоимость: %s руб.\n", state.Price))
			sb.WriteString(fmt.Sprintf("Дата сделки: %s\n", state.Date))
			sb.WriteString(fmt.Sprintf("Город сделки: %s\n", state.City))
			delete(userStates, userID)
			return c.Send(sb.String())
		default:
			delete(userStates, userID)
			return c.Send("Начнем заново. Введите ФИО продавца:")
		}
	})

	b.Handle("/help", func(c telebot.Context) error {
		helpMsg := "ℹ️ Я — бот для оформления договора купли-продажи автомобиля.\n" +
			"/start — начать работу\n" +
			"/reset — начать оформление заново\n" +
			"/help — справка по командам."
		return c.Send(helpMsg)
	})

	b.Handle("/reset", func(c telebot.Context) error {
		userID := c.Sender().ID
		delete(userStates, userID)
		return c.Send("Состояние сброшено. Введите 'Оформить договор' чтобы начать заново.")
	})
}
