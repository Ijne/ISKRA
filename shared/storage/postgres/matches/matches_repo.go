package matches

import (
	"database/sql"
	"iskra/shared/models"
)

type MatchesRepo struct {
	db *sql.DB
}

func New(db *sql.DB) (*MatchesRepo, error) {
	repo := MatchesRepo{db: db}
	err := repo.CreateTable()
	return &repo, err
}

func (r *MatchesRepo) CreateTable() error {
	// const op = "postgres.user.create_table"

	stmt := `
		CREATE TABLE IF NOT EXISTS public.matches (
			id int GENERATED ALWAYS AS IDENTITY NOT NULL,
			user1 int NOT NULL,
			user2 int NOT NULL,
			CONSTRAINT matches_pk PRIMARY KEY (id)
		);
	`

	_, err := r.db.Exec(stmt)
	return err
}

func (r *MatchesRepo) Exists(mothID int64, lightID int64) bool {
	stmt, err := r.db.Prepare(`
		SELECT EXISTS(
			SELECT 1 FROM matches
			WHERE moth_id = $1 AND light_id = $2 LIMIT 1
		)
	`)
	if err != nil {
		return false
	}

	var val bool
	err = stmt.QueryRow(mothID, lightID).Scan(&val)
	if err != nil {
		return false
	}
	return val
}

func (r *MatchesRepo) Create(match models.MatchDB) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO matches (user1, user2)
		VALUES ($1, $2)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(match.MothID, match.LightID)
	return err
}

func (r *MatchesRepo) Delete(mothID int64, lightID int64) error {
	stmt, err := r.db.Prepare(`
		DELETE FROM matches
		WHERE user1 = $1 AND user2 = $2
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(mothID, lightID)
	return err
}
