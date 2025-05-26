package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"carsbot/config"
	"carsbot/internal/bot"
	"carsbot/internal/fsm"
	"carsbot/internal/msg"
	"carsbot/internal/state"
)

func main() {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load config
	cfg, err := config.Load()
	if err != nil {
		logger.Error("config error", "err", err)
		panic(err)
	}

	// Initialize storage
	storage := state.NewInMemoryStorage()

	// Initialize FSM
	fsmEngine := fsm.NewFSM()

	// Initialize message generator
	messageGen := msg.NewMessageGenerator()

	// Initialize bot
	b := bot.NewBot(cfg, storage, fsmEngine, messageGen)

	// Start bot with graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go b.Start()

	<-ctx.Done()
	slog.Info("Bot stopped by signal (Ctrl+C or SIGTERM)")
}
