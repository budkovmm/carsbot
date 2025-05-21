package bot

import (
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
	return &Handler{storage: storage, fsm: fsm, msg: msg}
}

func (h *Handler) Register(b *telebot.Bot) {
	b.Handle("/start", func(c telebot.Context) error {
		return c.Send("👋 Здравствуйте! Я помогу вам оформить договор купли-продажи автомобиля.\n\nЧтобы начать, просто напишите любое сообщение или нажмите кнопку ниже.")
	})

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		userID := c.Sender().ID
		st, _ := h.storage.Get(userID)
		if st == nil {
			st = &state.UserState{Step: 0}
			h.storage.Set(userID, st)
		}
		resp, _ := h.fsm.NextStep(st, c.Text())
		h.storage.Set(userID, st)
		return c.Send(resp)
	})

	b.Handle("/reset", func(c telebot.Context) error {
		userID := c.Sender().ID
		h.storage.Delete(userID)
		return c.Send("Состояние сброшено. Введите 'Оформить договор' чтобы начать заново.")
	})

	b.Handle("/help", func(c telebot.Context) error {
		helpMsg := "ℹ️ Я — бот для оформления договора купли-продажи автомобиля.\n" +
			"/start — начать работу\n" +
			"/reset — начать оформление заново\n" +
			"/help — справка по командам."
		return c.Send(helpMsg)
	})
}
