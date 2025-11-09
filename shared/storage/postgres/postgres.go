package postgres

import (
	"database/sql"
	"fmt"
	"iskra/shared/config"
	"iskra/shared/storage/postgres/flames"
	"iskra/shared/storage/postgres/matches"
	"iskra/shared/storage/postgres/user"
	"iskra/shared/storage/repos"

	_ "github.com/lib/pq"
)

type Storage struct {
	db          *sql.DB
	UserRepo    repos.UserRepo
	MatchesRepo repos.MatchesRepo
	FlamesRepo  repos.FlamesRepo
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	db, err := sql.Open("postgres", getDSN(cfg))
	if err != nil {
		return nil, err
	}

	userRepo, err := user.New(db)
	if err != nil {
		return nil, fmt.Errorf("starting user repo: %w", err)
	}
	matchesRepo, err := matches.New(db)
	if err != nil {
		return nil, fmt.Errorf("starting matches repo: %w", err)
	}
	flamesRepo, err := flames.New(db)
	if err != nil {
		return nil, fmt.Errorf("starting flames repo: %w", err)
	}

	return &Storage{
		db:          db,
		UserRepo:    userRepo,
		MatchesRepo: matchesRepo,
		FlamesRepo:  flamesRepo,
	}, nil
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
