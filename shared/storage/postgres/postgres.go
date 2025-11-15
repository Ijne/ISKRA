package postgres

import (
	"database/sql"
	"fmt"
	"iskra/shared/config"
	"iskra/shared/storage/postgres/flames"
	"iskra/shared/storage/postgres/like"
	"iskra/shared/storage/postgres/matches"
	"iskra/shared/storage/postgres/user"
	"iskra/shared/storage/repos"

	_ "github.com/lib/pq"
)

type Storage struct {
	db          *sql.DB
	FlamesRepo  repos.FlamesRepo
	UserRepo    repos.UserRepo
	LikeRepo    repos.LikeRepo
	MatchesRepo repos.MatchesRepo
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

	likeRepo, err := like.New(db)
	if err != nil {
		return nil, fmt.Errorf("starting like repo: %w", err)
	}

	return &Storage{db: db, FlamesRepo: flamesRepo, UserRepo: userRepo, LikeRepo: likeRepo, MatchesRepo: matchesRepo}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
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
