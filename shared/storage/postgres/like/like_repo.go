package like

import (
	"database/sql"
	"iskra/shared/storage/repos"
)

type LikeRepo struct {
	db *sql.DB
}

func New(db *sql.DB) (repos.LikeRepo, error) {
	repo := LikeRepo{db: db}
	err := repo.CreateTable()
	return &repo, err
}

func (r *LikeRepo) CreateTable() error {
	// const op = "postgres.user.create_table"

	stmt := `
		CREATE TABLE IF NOT EXISTS likes (
			id integer NOT NULL,
			user1 integer NOT NULL,
			user2 integer NOT NULL
		);
	`

	_, err := r.db.Exec(stmt)
	return err
}

func (r *LikeRepo) SetLike(id1, id2 int64) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO likes (
			user1,
			user2
		)
		VALUES ($1, $2)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id1, id2)
	return err
}

func (r *LikeRepo) IsLike(id1, id2 int64) bool {
	stmt, err := r.db.Prepare(`
		SELECT id FROM likes WHERE user1 = $1 AND user2 = $2
	`)

	if err != nil {
		return false
	}

	var id int64
	if err := stmt.QueryRow(id1, id2).Scan(&id); err != nil {
		return false
	}

	return true
}

func (r *LikeRepo) GetAllLikesToUser(id int64) ([]int64, error) {
	return []int64{}, nil
}

func (r *LikeRepo) GetAllLikesFromUser(id int64) ([]int64, error) {
	return []int64{}, nil
}
