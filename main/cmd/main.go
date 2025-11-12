package main

import (
	"fmt"
	"iskra/bot"
	"iskra/miniapp"
	"iskra/shared/config"
	"iskra/shared/storage/memgraph"
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

	g, err := memgraph.NewStorage(cfg)
	if err != nil {
		fmt.Println("Memgraph is not started")
		// panic(err)
	}

	// bot
	b := bot.New(cfg)
	botDone := b.ListenUpdates()

	// miniapp
	app := miniapp.New(cfg, s, g, b)
	appDone := app.Listen()

	<-botDone
	<-appDone
}
