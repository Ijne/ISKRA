package main

import (
	"fmt"
	"iskra/bot"
	"iskra/miniapp"
	"iskra/shared/config"
	"iskra/shared/storage/memgraph"
	"iskra/shared/storage/postgres"
	"os"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/local.yaml"
	}

	// config
	cfg, err := config.New(configPath)
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
