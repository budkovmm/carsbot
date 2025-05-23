package bot

import (
	"log/slog"

	"carsbot/internal/state"
)

type MessageGenerator struct{}

func NewMessageGenerator(cfg interface{}) *MessageGenerator {
	slog.Info("message generator created")
	return &MessageGenerator{}
}

func (mg *MessageGenerator) GenerateByState(st *state.UserState) string {
	if st.Step == 1 {
		return "Введите ФИО продавца:"
	} else if st.Step == 2 {
		return "Введите ФИО покупателя:"
	} else if st.Step == 3 {
		return "Введите VIN автомобиля:"
	} else if st.Step == 4 {
		return "Введите марку и модель автомобиля:"
	} else if st.Step == 5 {
		return "Введите год выпуска автомобиля:"
	} else if st.Step == 6 {
		return "Введите цвет автомобиля:"
	} else if st.Step == 7 {
		return "Введите цену автомобиля:"
	} else if st.Step == 8 {
		return "Введите дату продажи:"
	} else if st.Step == 9 {
		return "Введите город продажи:"
	} else if st.Step == 10 {
		return "Спасибо! Вот все данные, которые вы ввели:\n" +
			"ФИО продавца: " + st.SellerName + "\n" +
			"ФИО покупателя: " + st.BuyerName + "\n" +
			"VIN: " + st.VIN + "\n" +
			"Марка/модель: " + st.BrandModel + "\n" +
			"Год выпуска: " + st.Year + "\n" +
			"Цвет: " + st.Color + "\n" +
			"Стоимость: " + st.Price + " руб.\n" +
			"Дата сделки: " + st.Date + "\n" +
			"Город сделки: " + st.City + "\n"
	}
	return ""
}

func (mg *MessageGenerator) Help() string {
	return "ℹ️ Я — бот для оформления договора купли-продажи автомобиля.\n" +
		"/start — начать работу\n" +
		"/reset — начать оформление заново\n" +
		"/help — справка по командам."
}

func (mg *MessageGenerator) Reset() string {
	return "Состояние сброшено. Введите 'Оформить договор' чтобы начать заново."
}

func (mg *MessageGenerator) Start() string {
	return "👋 Здравствуйте! Я помогу вам оформить договор купли-продажи автомобиля.\n\nЧтобы начать, просто напишите любое сообщение или нажмите кнопку ниже."
}
