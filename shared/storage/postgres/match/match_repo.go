package match

import (
	"database/sql"
	"iskra/shared/storage/repos"
)

type MatchRepo struct {
	db *sql.DB
}

func New(db *sql.DB) (repos.MatchRepo, error) {
	repo := MatchRepo{db: db}
	err := repo.CreateTable()
	return &repo, err
}

func (r *MatchRepo) CreateTable() error {
	// const op = "postgres.user.create_table"

	stmt := `
		CREATE TABLE IF NOT EXISTS public.matches (
			id int GENERATED ALWAYS AS IDENTITY NOT NULL,
			user1 int NOT NULL,
			user2 int NOT NULL,
			CONSTRAINT matches_pk PRIMARY KEY (id)
		);
	`
	// Дописать недостающее ^^^^^^

	_, err := r.db.Exec(stmt)
	return err
}

func (r *MatchRepo) SetMatch(id1, id2 int64) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO matches (
			user1,
			user2
		)
		VALUES ($1, $2)
	`)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(id1, id2); err != nil {
		return err
	}
	_, err = stmt.Exec(id2, id1)
	return err
}
