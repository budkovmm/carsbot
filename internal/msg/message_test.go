package msg

import (
	"carsbot/internal/state"
	"testing"
)

func TestMessageGenerator_Welcome(t *testing.T) {
	mg := &MessageGenerator{}
	msg := mg.Welcome()
	if msg == "" {
		t.Error("Welcome message should not be empty")
	}
}

func TestMessageGenerator_Start(t *testing.T) {
	mg := &MessageGenerator{}
	msg := mg.Start()
	if msg == "" {
		t.Error("Start message should not be empty")
	}
}

func TestMessageGenerator_Help(t *testing.T) {
	mg := &MessageGenerator{}
	help := mg.Help()
	if help == "" {
		t.Error("Help message should not be empty")
	}
}

func TestMessageGenerator_Reset(t *testing.T) {
	mg := &MessageGenerator{}
	help := mg.Reset()
	if help == "" {
		t.Error("Reset message should not be empty")
	}
}

func TestMessageGenerator_Error(t *testing.T) {
	mg := &MessageGenerator{}
	help := mg.Error()
	if help == "" {
		t.Error("Error message should not be empty")
	}
}

func TestMessageGenerator_ForStep(t *testing.T) {
	mg := &MessageGenerator{}
	steps := []struct {
		step   int
		expect string
	}{
		{1, "Введите ФИО продавца:"},
		{2, "Введите ФИО покупателя:"},
		{3, "Введите VIN автомобиля:"},
		{4, "Введите марку и модель автомобиля:"},
		{5, "Введите год выпуска автомобиля:"},
		{6, "Введите цвет автомобиля:"},
		{7, "Введите стоимость автомобиля (в рублях):"},
		{8, "Введите дату сделки (например, 12.06.2024):"},
		{9, "Введите город сделки:"},
		{10, "Спасибо! Договор оформлен."},
	}
	for _, s := range steps {
		st := &state.UserState{Step: s.step}
		msg := mg.ForStep(st)
		if msg != s.expect {
			t.Errorf("ForStep(%d): got %q, want %q", s.step, msg, s.expect)
		}
	}
	// Проверка для невалидного шага
	st := &state.UserState{Step: 99}
	msg := mg.ForStep(st)
	if msg == "" {
		t.Error("ForStep for unknown step should not be empty")
	}
}
