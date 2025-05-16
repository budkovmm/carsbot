package main

import (
	"log"
	"os"
	"time"

	"carsbot/interanal/bot"

	"gopkg.in/telebot.v4"
)

func setBotCommands(b *telebot.Bot) error {
	commands := []telebot.Command{
		{Text: "start", Description: "Начать работу с ботом"},
		{Text: "help", Description: "Получить справку"},
		{Text: "reset", Description: "Начать оформление заново"},
	}
	return b.SetCommands(commands)
}

func main() {
	pref := telebot.Settings{
		Token:  os.Getenv("TG_BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	if err := setBotCommands(b); err != nil {
		log.Println("setting bot commands failed", err)
	}

	bot.RegisterHandlers(b)
	b.Start()
}
