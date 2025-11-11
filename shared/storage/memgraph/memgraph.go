package memgraph

import (
	"fmt"
	"iskra/shared/config"
	"iskra/shared/storage/memgraph/socialweb"
	"iskra/shared/storage/repos"

	_ "github.com/lib/pq"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Storage struct {
	driver        *neo4j.DriverWithContext
	SocialWebRepo repos.SocialWebRepo
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	target, auth := getDSN(cfg)
	driver, err := neo4j.NewDriverWithContext(target, auth)
	if err != nil {
		return nil, err
	}

	socialWebRepo, err := socialweb.New(driver)
	if err != nil {
		return nil, err
	}

	return &Storage{driver: &driver, SocialWebRepo: socialWebRepo}, nil
}

func getDSN(cfg *config.Config) (string, neo4j.AuthToken) {
	return fmt.Sprintf(
		"%s://%s:%s",
		cfg.Memgraph.Protocol,
		cfg.Memgraph.Host,
		cfg.Memgraph.Port,
	), neo4j.BasicAuth(cfg.Memgraph.Username, cfg.Memgraph.Password, "")
}
