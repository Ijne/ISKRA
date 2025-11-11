package bot

import (
	"context"
	"fmt"
	"iskra/shared/config"
	"os"
	"os/signal"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

type Bot struct {
	api *maxbot.Api
	ctx context.Context
}

func New(cfg *config.Config) *Bot {
	api, err := maxbot.New(cfg.Bot.Token)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, os.Kill, os.Interrupt)
		<-exit
		cancel()
	}()

	bot := Bot{
		api: api,
		ctx: ctx,
	}

	return &bot
}

func (b *Bot) ListenUpdates() chan bool {
	done := make(chan bool)
	go func() {
		for upd := range b.api.GetUpdates(b.ctx) { // Чтение из канала с обновлениями
			switch upd := upd.(type) { // Определение типа пришедшего обновления
			case *schemes.MessageCreatedUpdate:
				// Отправка сообщения
				err := b.SendMessage(upd.Message.Recipient.ChatId, fmt.Sprintf("Hello from Bot! Your chat id: %d.", upd.Message.Recipient.ChatId))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		done <- true
	}()
	return done
}

func (b *Bot) SendMessage(chatID int64, msg string) error {
	_, err := b.api.Messages.Send(b.ctx,
		maxbot.NewMessage().
			SetChat(chatID).
			SetText(msg))
	return err
}

func (b *Bot) SendNick(chatID int64, nick string) {
	fmt.Println("send nick")
	msg := "ℹ️ Вас лайкнул(а): " + nick + "."
	b.SendMessage(chatID, msg)
}
