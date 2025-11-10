// package main

// import "fmt"

// func main() {
// 	fmt.Println("Here i am")

// 	// TODO: config loading

// 	// TODO: logger

// 	// TODO: storage

// 	// TODO: router

// 	// TODO: listen http
// }

package main

import (
	"context"
	"fmt"
	"iskra/shared/config"
	"os"
	"os/signal"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

func main() {
	// config
	cfg, err := config.New("./config/local.yaml")
	if err != nil {
		panic(err)
	}

	api, err := maxbot.New(cfg.Bot.Token)
	if err != nil {
		panic(err)
	}

	// Some methods demo:
	ctx := context.TODO()
	info, err := api.Bots.GetBot(ctx)
	fmt.Printf("Get me: %#v %#v", info, err)

	ctx, cancel := context.WithCancel(context.Background()) // создам
	go func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, os.Kill, os.Interrupt)
		<-exit
		cancel()
	}()

	for upd := range api.GetUpdates(ctx) { // Чтение из канала с обновлениями
		switch upd := upd.(type) { // Определение типа пришедшего обновления
		case *schemes.MessageCreatedUpdate:
			// Отправка сообщения
			_, err := api.Messages.Send(ctx, maxbot.NewMessage().SetChat(upd.Message.Recipient.ChatId).SetText("Hello from Bot"))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
