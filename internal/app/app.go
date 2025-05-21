package app

import (
	"carsbot/config"
	"carsbot/internal/bot"
	"carsbot/internal/fsm"
	"carsbot/internal/state"
)

func NewApp(cfg *config.Config) *bot.Bot {
	storage := state.NewInMemoryStorage()
	fsmEngine := fsm.NewFSM()
	messageGen := bot.NewMessageGenerator(cfg)
	return bot.NewBot(cfg, storage, fsmEngine, messageGen)
}
