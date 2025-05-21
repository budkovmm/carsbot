package fsm

import "carsbot/internal/state"

type FSM interface {
	NextStep(st *state.UserState, input string) (string, error)
}

type SimpleFSM struct{}

func NewFSM() *SimpleFSM { return &SimpleFSM{} }

func (f *SimpleFSM) NextStep(st *state.UserState, input string) (string, error) {
	switch st.Step {
	case 0:
		st.Step = 1
		return "Введите ФИО продавца:", nil
	case 1:
		st.SellerName = input
		st.Step = 2
		return "Введите ФИО покупателя:", nil
	case 2:
		st.BuyerName = input
		st.Step = 3
		return "Введите VIN автомобиля:", nil
	case 3:
		st.VIN = input
		st.Step = 4
		return "Введите марку и модель автомобиля:", nil
	case 4:
		st.BrandModel = input
		st.Step = 5
		return "Введите год выпуска автомобиля:", nil
	case 5:
		st.Year = input
		st.Step = 6
		return "Введите цвет автомобиля:", nil
	case 6:
		st.Color = input
		st.Step = 7
		return "Введите стоимость автомобиля (в рублях):", nil
	case 7:
		st.Price = input
		st.Step = 8
		return "Введите дату сделки (например, 12.06.2024):", nil
	case 8:
		st.Date = input
		st.Step = 9
		return "Введите город сделки:", nil
	case 9:
		st.City = input
		st.Step = 10
		return "Спасибо! Вот все данные, которые вы ввели:\n" +
			"ФИО продавца: " + st.SellerName + "\n" +
			"ФИО покупателя: " + st.BuyerName + "\n" +
			"VIN: " + st.VIN + "\n" +
			"Марка/модель: " + st.BrandModel + "\n" +
			"Год выпуска: " + st.Year + "\n" +
			"Цвет: " + st.Color + "\n" +
			"Стоимость: " + st.Price + " руб.\n" +
			"Дата сделки: " + st.Date + "\n" +
			"Город сделки: " + st.City + "\n", nil
	default:
		return "Оформление завершено. Введите /reset для нового договора.", nil
	}
}
