package user

import (
	"database/sql"
	"iskra/shared/models"
)

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) (*UserRepo, error) {
	repo := UserRepo{db: db}
	err := repo.CreateTable()
	return &repo, err
}

func (r *UserRepo) CreateTable() error {
	// const op = "postgres.user.create_table"

	stmt := `
		CREATE TABLE IF NOT EXISTS users (
			chat_id INTEGER NOT NULL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			nick VARCHAR(255) NOT NULL,
			description TEXT,
			icon VARCHAR(255)
		)
	`

	_, err := r.db.Exec(stmt)
	return err
}

func (r *UserRepo) GetUser(chatID int64) (models.UserDB, error) {
	stmt, err := r.db.Prepare(`
		SELECT chat_id, username, nick, description, icon
		FROM users
		WHERE chat_id = $1
	`)

	res := models.UserDB{}
	if err != nil {
		return res, err
	}

	err = stmt.QueryRow(chatID).Scan(&res.ChatID, &res.Username, &res.Nick, &res.Description, &res.Icon)
	return res, err
}

func (r *UserRepo) CreateUser(user models.UserCreate) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO users (
			chat_id,
			username,
			nick,
			description,
			icon
		)
		VALUES ($1, $2, $3, $4, $5)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.ChatID, user.Username, user.Nick, user.Description, user.Icon)
	return err
}
