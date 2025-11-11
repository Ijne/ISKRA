package main

import (
	"iskra/bot"
	"iskra/miniapp"
	"iskra/shared/config"
	"iskra/shared/storage/postgres"
)

func main() {
	cfg, err := config.New("./config/local.yaml")
	if err != nil {
		panic(err)
	}

	// storage
	s, err := postgres.NewStorage(cfg)
	if err != nil {
		panic(err)
	}

	// bot
	b := bot.New(cfg)
	botDone := b.ListenUpdates()

	// miniapp
	app := miniapp.New(cfg, s, b)
	appDone := app.Listen()

	<-botDone
	<-appDone
}
