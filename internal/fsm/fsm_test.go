package fsm

import (
	"carsbot/internal/state"
	"testing"
)

func TestFSM_Transition(t *testing.T) {
	fsm := NewFSM()
	st := &state.UserState{}

	inputs := []struct {
		input  string
		setter func(*state.UserState) string // функция для получения значения из state
		step   int
	}{
		{"Иванов Иван", func(st *state.UserState) string { return st.SellerName }, 2},
		{"Петров Петр", func(st *state.UserState) string { return st.BuyerName }, 3},
		{"VIN123", func(st *state.UserState) string { return st.VIN }, 4},
		{"Toyota Camry", func(st *state.UserState) string { return st.BrandModel }, 5},
		{"2018", func(st *state.UserState) string { return st.Year }, 6},
		{"Белый", func(st *state.UserState) string { return st.Color }, 7},
		{"850000", func(st *state.UserState) string { return st.Price }, 8},
		{"12.06.2024", func(st *state.UserState) string { return st.Date }, 9},
		{"Москва", func(st *state.UserState) string { return st.City }, 10},
	}

	// Первый переход (шаг 0 -> 1)
	fsm.Transition(st, "")
	if st.Step != 1 {
		t.Errorf("expected step 1 after first transition, got %d", st.Step)
	}

	for i, test := range inputs {
		fsm.Transition(st, test.input)
		if st.Step != test.step {
			t.Errorf("step %d: expected step %d, got %d", i, test.step, st.Step)
		}
		if v := test.setter(st); v != test.input {
			t.Errorf("step %d: expected value %q, got %q", i, test.input, v)
		}
	}

	// Проверяем, что после последнего шага step не увеличивается
	fsm.Transition(st, "что-то ещё")
	if st.Step != 11 {
		t.Errorf("after finish: expected step 11, got %d", st.Step)
	}
}
