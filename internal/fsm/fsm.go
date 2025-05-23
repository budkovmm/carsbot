package fsm

import (
	"log/slog"

	"carsbot/internal/state"
)

type FSM interface {
	Transition(st *state.UserState, input string)
}

type SimpleFSM struct{}

func NewFSM() *SimpleFSM {
	slog.Info("fsm created")
	return &SimpleFSM{}
}

func (f *SimpleFSM) Transition(st *state.UserState, input string) {
	switch st.Step {
	case 0:
		st.Step = 1
		return
	case 1:
		st.SellerName = input
		st.Step = 2
		return
	case 2:
		st.BuyerName = input
		st.Step = 3
		return
	case 3:
		st.VIN = input
		st.Step = 4
		return
	case 4:
		st.BrandModel = input
		st.Step = 5
		return
	case 5:
		st.Year = input
		st.Step = 6
		return
	case 6:
		st.Color = input
		st.Step = 7
		return
	case 7:
		st.Price = input
		st.Step = 8
		return
	case 8:
		st.Date = input
		st.Step = 9
		return
	case 9:
		st.City = input
		st.Step = 10
		return
	case 10:
		st.Step = 11
		return
	default:
		return
	}
}
