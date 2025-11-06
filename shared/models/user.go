package models

// что будет возвращаться из postgres
type UserDB struct {
	ChatID      int64
	Username    string // мб потом уберём
	Nick        string // как будет отображаться в приложении, по умолчанию будет составляться из ФИ
	Description string
	Icon        string // url изображения
}

type UserCreate struct {
	ChatID      int64
	Username    string // мб потом уберём
	Nick        string // как будет отображаться в приложении, по умолчанию будет составляться из ФИ
	Description string
	Icon        string // url изображения
}

type UserResponse struct {
}

type UserRequest struct {
}
