package postgres

import (
	"database/sql"
	"fmt"
	"iskra/shared/config"
	"iskra/shared/storage/postgres/like"
	"iskra/shared/storage/postgres/match"
	"iskra/shared/storage/postgres/user"
	"iskra/shared/storage/repos"

	_ "github.com/lib/pq"
)

type Storage struct {
	db        *sql.DB
	UserRepo  repos.UserRepo
	LikeRepo  repos.LikeRepo
	MatchRepo repos.MatchRepo
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	db, err := sql.Open("postgres", getDSN(cfg))
	if err != nil {
		return nil, err
	}

	userRepo, err := user.New(db)
	if err != nil {
		return nil, err
	}

	likeRepo, err := like.New(db)
	if err != nil {
		return nil, err
	}

	matchRepo, err := match.New(db)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db, UserRepo: userRepo, LikeRepo: likeRepo, MatchRepo: matchRepo}, nil
}

func getDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
		cfg.Postgres.SSLMode,
	)
}
