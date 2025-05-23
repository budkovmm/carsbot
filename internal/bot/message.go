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

func (mg *MessageGenerator) Welcome() string {
	return "👋 Здравствуйте! Я помогу вам оформить договор купли-продажи автомобиля.\n\nЧтобы начать, просто напишите любое сообщение или нажмите кнопку ниже."
}

func (mg *MessageGenerator) ForStep(st *state.UserState) string {
	switch st.Step {
	case 1:
		return "Введите ФИО продавца:"
	case 2:
		return "Введите ФИО покупателя:"
	case 3:
		return "Введите VIN автомобиля:"
	case 4:
		return "Введите марку и модель автомобиля:"
	case 5:
		return "Введите год выпуска автомобиля:"
	case 6:
		return "Введите цвет автомобиля:"
	case 7:
		return "Введите стоимость автомобиля (в рублях):"
	case 8:
		return "Введите дату сделки (например, 12.06.2024):"
	case 9:
		return "Введите город сделки:"
	case 10:
		return "Спасибо! Договор оформлен."
	default:
		return "Оформление завершено. Введите /reset для нового договора."
	}
}

func (mg *MessageGenerator) Help() string {
	return "ℹ️ Я — бот для оформления договора купли-продажи автомобиля.\n" +
		"/start — начать работу\n" +
		"/reset — начать оформление заново\n" +
		"/help — справка по командам."
}

func (mg *MessageGenerator) Reset() string {
	return "Состояние сброшено. Введите /start чтобы начать заново."
}

func (mg *MessageGenerator) Start() string {
	return "👋 Здравствуйте! Я помогу вам оформить договор купли-продажи автомобиля.\n\nЧтобы начать, просто напишите любое сообщение или нажмите кнопку ниже."
}

func (mg *MessageGenerator) Error() string {
	return "Произошла ошибка при получении состояния. Попробуйте снова."
}
