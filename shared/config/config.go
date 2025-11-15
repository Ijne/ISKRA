package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert/yaml"
)

type Config struct {
	Server struct {
		Address string `yaml:"address" validate:"required"`
	} `yaml:"server"`

	Bot struct {
		Token    string
		Username string `yaml:"username" validate:"required"`
	} `yaml:"bot"`

	MiniApp struct {
		Proto string `yaml:"proto"`
		Host  string `yaml:"host"`
		Port  string `yaml:"port"`
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

	JWT struct {
		SecretKey       []byte
		TokenTTLSeconds int `yaml:"token_ttl_seconds" validate:"required"`
	} `yaml:"jwt"`
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
		log.Println("cfg: no token for timepad")
	}
	cfg.JWT.SecretKey = []byte(os.Getenv("JWT_TOKEN"))
	if len(cfg.JWT.SecretKey) == 0 {
		cfg.JWT.SecretKey = []byte("secret_key")
	}
	cfg.Bot.Token = os.Getenv("BOT_TOKEN")
	if len(cfg.Bot.Token) == 0 {
		cfg.Bot.Token = "no token"
		log.Println("cfg: no token for bot")
	}

	cfg.Memgraph.Username = os.Getenv("MEMGRAPH_USER")
	if cfg.Memgraph.Username == "" {
		cfg.Memgraph.Username = ""
	}
	cfg.Memgraph.Password = os.Getenv("MEMGRAPH_PASSWORD")
	if cfg.Memgraph.Password == "" {
		cfg.Memgraph.Password = ""
	}

	// address for miniapp`s backend can be set using environment vars
	miniappProto := os.Getenv("MINIAPP_PROTO")
	if miniappProto != "" {
		cfg.MiniApp.Proto = miniappProto
	} else {
		log.Println("cfg: proto for miniapp extracted from yaml")
	}
	miniappHost := os.Getenv("MINIAPP_HOST")
	if miniappHost != "" {
		cfg.MiniApp.Host = miniappHost
	} else {
		log.Println("cfg: host for miniapp extracted from yaml")
	}
	miniappPort := os.Getenv("MINIAPP_PORT")
	if miniappPort != "" {
		cfg.MiniApp.Port = miniappPort
	} else {
		log.Println("cfg: port for miniapp extracted from yaml")
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress != "" {
		cfg.Server.Address = serverAddress
	} else {
		log.Println("cfg: address for server extracted from yaml")
	}

	return &cfg, nil
}
