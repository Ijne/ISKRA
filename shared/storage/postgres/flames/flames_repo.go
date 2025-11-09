package flames

import (
	"database/sql"
	"iskra/shared/models"
)

type FlamesRepo struct {
	db *sql.DB
}

func New(db *sql.DB) (*FlamesRepo, error) {
	repo := FlamesRepo{db: db}
	err := repo.CreateTable()
	return &repo, err
}

func (r *FlamesRepo) CreateTable() error {
	// const op = "postgres.user.create_table"

	stmt := `
		CREATE TABLE IF NOT EXISTS public.flames (
			event_id int NOT NULL,
			user_id int NOT NULL,
			description TEXT,
			PRIMARY KEY (event_id, user_id)
		);

		CREATE TABLE IF NOT EXISTS public.events_categories (
			id int NOT NULL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS public.saved_events_categories (
			event_id int NOT NULL,
			category_id int NOT NULL
		);

		CREATE TABLE IF NOT EXISTS public.saved_events (
			id int NOT NULL PRIMARY KEY,
			starts_at TIMESTAMP NOT NULL,
			name VARCHAR NOT NULL,
			url VARCHAR NOT NULL,
			photo VARCHAR NOT NULL
		);
	`

	_, err := r.db.Exec(stmt)
	return err
}

func (r *FlamesRepo) Create(flame models.FlameDB) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO flames (event_id, user_id, description)
		VALUES ($1, $2, $3)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(flame.EventID, flame.UserID, flame.Description)
	return err
}

func (r *FlamesRepo) Update(flame models.FlameDB) error {
	stmt, err := r.db.Prepare(`
		UPDATE flames SET description = $1
		WHERE event_id = $2 AND user_id = $3
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(flame.Description, flame.EventID, flame.UserID)
	return err
}

func (r *FlamesRepo) GetLim(limit int) ([]models.FlameDB, error) {
	stmt, err := r.db.Prepare(`
		SELECT event_id, user_id, description
		FROM flames
		LIMIT $1
	`)
	if err != nil {
		return nil, err
	}

	res := make([]models.FlameDB, 0, limit)

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var curr models.FlameDB
		rows.Scan(&curr.EventID, &curr.UserID, &curr.Description)
		res = append(res, curr)
	}

	return res, nil
}

func (r *FlamesRepo) CreateEvent(event models.EventDB) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO saved_events (id, starts_at, name, url, photo)
		VALUES ($1, $2, $3, $4, $5)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.ID, event.StartsAt, event.Name, event.Url, event.Photo)
	return err
}

func (r *FlamesRepo) GetEvents() ([]models.EventDB, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, starts_at, name, url, photo
		FROM saved_events
	`)
	if err != nil {
		return nil, err
	}

	res := make([]models.EventDB, 0)

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var curr models.EventDB
		rows.Scan(&curr.ID, &curr.StartsAt, &curr.Name, &curr.Url, &curr.Photo)
		res = append(res, curr)
	}

	return res, nil
}

func (r *FlamesRepo) FillCategories(categories []models.EventCategory) error {
	for _, cat := range categories {
		stmt, err := r.db.Prepare(`
			INSERT INTO events_categories (id, name)
			VALUES ($1, $2)
		`)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(cat.ID, cat.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *FlamesRepo) GetByEvent(eventID int64) ([]models.FlameDB, error) {
	stmt, err := r.db.Prepare(`
		SELECT event_id, user_id, description
		FROM flames
		WHERE event_id = $1
	`)
	if err != nil {
		return nil, err
	}

	res := make([]models.FlameDB, 0)

	rows, err := stmt.Query(eventID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var curr models.FlameDB
		rows.Scan(&curr.EventID, &curr.UserID, &curr.Description)
		res = append(res, curr)
	}

	return res, nil
}

func (r *FlamesRepo) GetByUser(userID int64) ([]models.FlameDB, error) {
	stmt, err := r.db.Prepare(`
		SELECT event_id, user_id, description
		FROM flames
		WHERE user_id = $1
	`)
	if err != nil {
		return nil, err
	}

	res := make([]models.FlameDB, 0)

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var curr models.FlameDB
		rows.Scan(&curr.EventID, &curr.UserID, &curr.Description)
		res = append(res, curr)
	}

	return res, nil
}

func (r *FlamesRepo) Delete(eventID int64, userID int64) error {
	stmt, err := r.db.Prepare(`
		DELETE FROM flames
		WHERE event_id = $1 AND user_id = $2
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(eventID, userID)
	return err
}
