// package bot

// import (
// 	"context"
// 	"fmt"
// 	"iskra/shared/config"
// 	"log"
// 	"os"
// 	"os/signal"

// 	maxbot "github.com/max-messenger/max-bot-api-client-go"
// 	"github.com/max-messenger/max-bot-api-client-go/schemes"
// )

// type Bot struct {
// 	api *maxbot.Api
// 	ctx context.Context
// }

// func New(cfg *config.Config) *Bot {
// 	api, err := maxbot.New(cfg.Bot.Token)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func() {
// 		exit := make(chan os.Signal)
// 		signal.Notify(exit, os.Kill, os.Interrupt)
// 		<-exit
// 		cancel()
// 	}()

// 	bot := Bot{
// 		api: api,
// 		ctx: ctx,
// 	}

// 	return &bot
// }

// func (b *Bot) ListenUpdates() chan bool {
// 	done := make(chan bool)
// 	go func() {
// 		for upd := range b.api.GetUpdates(b.ctx) { // Чтение из канала с обновлениями
// 			log.Printf("***** %v\n", upd)
// 			log.Printf("***** %v\n", upd.GetChatID())
// 			log.Printf("***** %v\n", upd.GetUpdateType())
// 			switch upd := upd.(type) { // Определение типа пришедшего обновления
// 			case *schemes.MessageCreatedUpdate:
// 				// Отправка сообщения
// 				// err := b.SendMessage(upd.Message.Recipient.ChatId, fmt.Sprintf("Hello from Bot! Your chat id: %d.", upd.Message.Recipient.ChatId))
// 				err := b.SendMessage(upd.Message.Recipient.ChatId, fmt.Sprintf("Hello from Bot!\nYour chat id: %d;\nyour user id: %d;\nyour nick: %s;\nmessage: %s.",
// 					upd.Message.Recipient.ChatId,
// 					upd.Message.Recipient.UserId,
// 					upd.Message.Sender.Username,
// 					upd.Message.Body.Text,
// 				))
// 				if err != nil {
// 					// fmt.Println(err)
// 				}
// 			}
// 		}
// 		done <- true
// 	}()
// 	return done
// }

// func (b *Bot) SendMessage(chatID int64, msg string) error {
// 	_, err := b.api.Messages.Send(b.ctx,
// 		maxbot.NewMessage().
// 			SetChat(chatID).
// 			SetText(msg))
// 	return err
// }

// func (b *Bot) SendNick(chatID int64, nick string) {
// 	fmt.Println("send nick")
// 	msg := "ℹ️ Вас лайкнул(а): " + nick + "."
// 	b.SendMessage(chatID, msg)
// }

// package bot

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"iskra/shared/config"
// 	"net/http"
// 	"time"

// 	maxbot "github.com/max-messenger/max-bot-api-client-go"
// )

// type Bot struct {
// 	token     string
// 	apiClient *http.Client
// 	apiUrl    string
// 	marker    *int64
// 	ctx       context.Context
// }

// func New(cfg *config.Config) *Bot {
// 	req, err := http.NewRequest("GET", fullURL, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "Bearer "+t.token)
// 	resp, err := t.client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("error in request to timepad: %w", err)
// 	}

// 	api, err := maxbot.New(cfg.Bot.Token)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ctx, _ := context.WithCancel(context.Background())

// 	bot := Bot{
// 		token:     cfg.Bot.Token,
// 		apiClient: &http.Client{},
// 		apiUrl:    "https://platform-api.max.ru/",
// 		ctx:       ctx,
// 	}

// 	return &bot
// }

// func (*Bot) ListenUpdates() <-chan bool {
// 	ch := make(chan bool)
// 	go func() {

// 		ch <- true
// 	}()
// 	return ch
// }

package bot

import (
	"encoding/json"
	"fmt"
	"iskra/bot/internal/models"
	"log"
	"net/http"
	"time"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
)

type MaxBot struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	Marker     *int64
}

func New(token string) *MaxBot {
	return &MaxBot{
		BaseURL: "https://platform-api.max.ru",
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: 40 * time.Second,
		},
	}
}

func (b *MaxBot) GetUpdates(timeout int) (*models.UpdatesResponse, error) {
	url := fmt.Sprintf("%s/updates", c.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %v", err)
	}

	q := req.URL.Query()
	// q.Add("access_token", c.Token)
	req.Header.Add("Authorization", "Bearer "+b.Token)
	q.Add("timeout", fmt.Sprintf("%d", timeout))
	q.Add("limit", "100")

	if b.Marker != nil {
		q.Add("marker", fmt.Sprintf("%d", *b.Marker))
	}

	req.URL.RawQuery = q.Encode()

	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка API: статус %d", resp.StatusCode)
	}

	var updatesResp models.UpdatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&updatesResp); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	// сохраняем маркер для следующего запроса
	if updatesResp.Marker != nil {
		b.Marker = updatesResp.Marker
	}

	return &updatesResp, nil
}

func (b *MaxBot) ListenUpdates() chan bool {
	done := make(chan bool)
	go func() {
		for upd, err := range b.GetUpdates(30) {
			log.Printf("***** %v\n", upd)
			log.Printf("***** %v\n", upd.GetChatID())
			log.Printf("***** %v\n", upd.GetUpdateType())

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
