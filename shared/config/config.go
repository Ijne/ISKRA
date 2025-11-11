package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert/yaml"
)

type Config struct {
	Bot struct {
		Host string `yaml:"host" validate:"required"`
		Port string `yaml:"port" validate:"required"`
	} `yaml:"bot"`

	MiniApp struct {
		Host          string `yaml:"host" validate:"required"`
		Port          string `yaml:"port" validate:"required"`
		StaticPath    string `yaml:"static-path" validate:"required"`
		TemplatesPath string `yaml:"templates-path" validate:"required"`
	} `yaml:"mini-app"`

	Postgres struct {
		Host     string `yaml:"host" validate:"required"`
		Port     string `yaml:"port" validate:"required"`
		Username string
		Password string
		Database string `yaml:"database" validate:"required"`
		SSLMode  string `yaml:"ssl_mode" validate:"required"`
	} `yaml:"postgres"`

	Memgraph struct {
		Protocol string `yaml:"protocol" validate:"required"`
		Host     string `yaml:"host" validate:"required"`
		Port     string `yaml:"port" validate:"required"`
		Username string
		Password string
	} `yaml:"memgraph"`

	Timepad struct {
		Token string
	}
}

func New(configPath string) (*Config, error) {
	workDir, _ := os.Getwd()
	log.Printf("WD for loading config: %s\n", workDir)

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// чтение yaml
	var cfg Config
	yaml.Unmarshal(file, &cfg)
	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return nil, err
	}

	// чтения переменных окружения
	cfg.Postgres.Username = os.Getenv("POSTGRES_USER")
	if cfg.Postgres.Username == "" {
		cfg.Postgres.Username = "postgres"
	}
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	if cfg.Postgres.Password == "" {
		cfg.Postgres.Password = "postgres"
	}
	cfg.Timepad.Token = os.Getenv("TIMEPAD_TOKEN")
	if cfg.Timepad.Token == "" {
		cfg.Timepad.Token = "no token"
		fmt.Println("cfg: no token for timepad")
	}

	cfg.Memgraph.Username = os.Getenv("MEMGRAPH_USER")
	if cfg.Memgraph.Username == "" {
		cfg.Memgraph.Username = ""
	}
	cfg.Memgraph.Password = os.Getenv("MEMGRAPH_PASSWORD")
	if cfg.Memgraph.Password == "" {
		cfg.Memgraph.Password = ""
	}

	return &cfg, nil
}
