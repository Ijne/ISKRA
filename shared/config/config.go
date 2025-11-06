package config

import (
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
}

func New(configPath string) (*Config, error) {
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
	// ...

	return &cfg, nil
}
