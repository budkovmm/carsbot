package bot

import (
	"log/slog"

	"carsbot/internal/fsm"
	"carsbot/internal/state"

	"gopkg.in/telebot.v4"
)

type Handler struct {
	storage state.StateStorage
	fsm     fsm.FSM
	msg     *MessageGenerator
}

func NewHandler(storage state.StateStorage, fsm fsm.FSM, msg *MessageGenerator) *Handler {
	slog.Info("handler created")
	return &Handler{storage: storage, fsm: fsm, msg: msg}
}

func (h *Handler) Register(b *telebot.Bot) {
	b.Handle("/start", func(c telebot.Context) error {
		return c.Send(h.msg.Start())
	})

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		userID := c.Sender().ID
		st, _ := h.storage.Get(userID)
		if st == nil {
			st = &state.UserState{Step: 0}
			h.storage.Set(userID, st)
		}
		h.fsm.Transition(st, c.Text())
		h.storage.Set(userID, st)
		return c.Send(h.msg.GenerateByState(st))
	})

	b.Handle("/reset", func(c telebot.Context) error {
		userID := c.Sender().ID
		h.storage.Delete(userID)
		return c.Send(h.msg.Reset())
	})

	b.Handle("/help", func(c telebot.Context) error {
		return c.Send(h.msg.Help())
	})
}
