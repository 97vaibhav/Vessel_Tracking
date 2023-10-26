package utils

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug            bool   `default:"false" split_words:"true"`
	ServerAddress    string `required:"true" split_words:"true"`
	DatabaseHost     string `required:"true" split_words:"true"`
	DatabasePort     string `required:"true" split_words:"true"`
	DatabaseUser     string `required:"true" split_words:"true"`
	DatabasePassword string `required:"true" split_words:"true"`
	DatabaseName     string `required:"true" split_words:"true"`
}

func LoadConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &c, nil
}
