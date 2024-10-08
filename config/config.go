package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"8080"`
}

func New() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
