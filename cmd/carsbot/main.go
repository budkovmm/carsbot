package main

import (
	"carsbot/config"
	"carsbot/internal/app"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}
	a := app.NewApp(cfg)
	a.Start()
}
