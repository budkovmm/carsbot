package bot

import (
	"carsbot/config"
	"carsbot/internal/fsm"
	"carsbot/internal/handler"
	"carsbot/internal/msg"
	"carsbot/internal/state"

	"log/slog"
	"time"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	bot     *telebot.Bot
	storage state.StateStorage
	fsm     fsm.FSM
	msg     *msg.MessageGenerator
	handler *handler.Handler
}

func LogHandlerDuration(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		user := c.Sender()
		start := time.Now()
		err := next(c)
		dur := time.Since(start)
		slog.Info("handler finished", "user_id", user.ID, "username", user.Username, "duration_ms", dur.Milliseconds(), "message", c.Text())
		return err
	}
}

func NewBot(cfg *config.Config, storage state.StateStorage, fsm fsm.FSM, msg *msg.MessageGenerator) *Bot {
	b, err := telebot.NewBot(telebot.Settings{Token: cfg.TelegramToken})
	if err != nil {
		slog.Error("failed to create bot", "err", err)
		panic(err)
	}
	botInstance := &Bot{
		bot:     b,
		storage: storage,
		fsm:     fsm,
		msg:     msg,
	}
	b.Use(LogHandlerDuration)
	botInstance.handler = handler.New(storage, fsm, msg)
	botInstance.handler.Register(b)
	return botInstance
}

func (b *Bot) Start() {
	slog.Info("bot started and polling for updates")
	b.bot.Start()
	slog.Info("Bot stopped")
}
