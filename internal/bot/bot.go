package bot

import (
	"carsbot/config"
	"carsbot/internal/fsm"
	"carsbot/internal/state"
	"log"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	bot     *telebot.Bot
	storage state.StateStorage
	fsm     fsm.FSM
	msg     *MessageGenerator
	handler *Handler
}

func NewBot(cfg *config.Config, storage state.StateStorage, fsm fsm.FSM, msg *MessageGenerator) *Bot {
	b, err := telebot.NewBot(telebot.Settings{Token: cfg.TelegramToken})
	if err != nil {
		log.Fatal(err)
	}
	botInstance := &Bot{
		bot:     b,
		storage: storage,
		fsm:     fsm,
		msg:     msg,
	}
	botInstance.handler = NewHandler(storage, fsm, msg)
	botInstance.handler.Register(b)
	return botInstance
}

func (b *Bot) Start() {
	b.bot.Start()
}
