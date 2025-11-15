package models

import (
	"encoding/json"
)

// Структуры для API ответов
type UpdatesResponse struct {
	Updates []Update `json:"updates"`
	Marker  *int64   `json:"marker,omitempty"`
}

type Update struct {
	UpdateType string    `json:"update_type"`
	Timestamp  int64     `json:"timestamp"`
	Callback   *Callback `json:"callback,omitempty"`
	Message    *Message  `json:"message,omitempty"`
	ChatID     *int64    `json:"chat_id,omitempty"`
	User       *User     `json:"user,omitempty"`
	UserLocale *string   `json:"user_locale,omitempty"`
	Payload    *string   `json:"payload,omitempty"`
}

type Callback struct {
	Timestamp  int64  `json:"timestamp"`
	CallbackID string `json:"callback_id"`
	Payload    string `json:"payload"`
	User       User   `json:"user"`
}

type User struct {
	UserID           int64   `json:"user_id"`
	FirstName        string  `json:"first_name"`
	LastName         *string `json:"last_name,omitempty"`
	Username         *string `json:"username,omitempty"`
	IsBot            bool    `json:"is_bot"`
	LastActivityTime int64   `json:"last_activity_time"`
}

type Message struct {
	Sender    User        `json:"sender"`
	Recipient Recipient   `json:"recipient"`
	Timestamp int64       `json:"timestamp"`
	Body      MessageBody `json:"body"`
}

type Recipient struct {
	ChatID   *int64 `json:"chat_id,omitempty"`
	ChatType string `json:"chat_type"`
	UserID   *int64 `json:"user_id,omitempty"`
}

type MessageBody struct {
	MID         string       `json:"mid"`
	Seq         int64        `json:"seq"`
	Text        *string      `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type InlineKeyboard struct {
	Buttons [][]Button `json:"buttons"`
}

type Button struct {
	Type   string  `json:"type"`
	Text   string  `json:"text"`
	WebApp *string `json:"web_app,omitempty"`
	// ContactID *int64  `json:"contact_id,omitempty"`
	Payload *string `json:"payload,omitempty"`
	// URL       *string `json:"url,omitempty"`
}

type OpenAppButtonData struct {
	WebApp     string `json:"web_app"`
	ContactID  int64  `json:"contact_id"`
	Payload    string `json:"payload"`
	UserID     int64  `json:"user_id"`
	CallbackID string `json:"callback_id"`
}

type MessageRequest struct {
	Text string `json:"text"`
}

type MessageOpenAppButtonRequest struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Buttons struct {
	Buttons [][]Button `json:"buttons"`
}

// {
//     "text": "Запустите мини-приложение:",
//     "attachments": [
//       {
//       "type": "inline_keyboard",
//       "payload": {
//         "buttons": [
//         [
//           {
//           "type": "open_app",
//           "text": "Запустить приложение",
//           "web_app": "t21_hakaton_bot",
//           "payload": "http___localhost___8080"
//           }
//         ]
//         ]
//       }
//       }
//     ]
//   }
