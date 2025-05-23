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
	b.Handle("/start", h.OnStart)
	b.Handle(telebot.OnText, h.OnText)
	b.Handle("/reset", h.OnReset)
	b.Handle("/help", h.OnHelp)
}

func (h *Handler) OnStart(c telebot.Context) error {
	user := c.Sender()
	slog.Info("start command received", "user_id", user.ID, "username", user.Username)
	return c.Send(h.msg.Welcome())
}

func (h *Handler) OnText(c telebot.Context) error {
	userID := c.Sender().ID
	st, err := h.storage.Get(userID)
	if err != nil {
		slog.Error("error getting state", "error", err)
		return c.Send(h.msg.Error())
	}
	if st == nil {
		st = &state.UserState{Step: 0}
		h.storage.Set(userID, st)
	}
	h.fsm.Transition(st, c.Text())
	h.storage.Set(userID, st)
	return c.Send(h.msg.ForStep(st))
}

func (h *Handler) OnReset(c telebot.Context) error {
	userID := c.Sender().ID
	h.storage.Delete(userID)
	return c.Send(h.msg.Reset())
}

func (h *Handler) OnHelp(c telebot.Context) error {
	return c.Send(h.msg.Help())
}
