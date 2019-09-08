package config

import "github.com/caarlos0/env/v6"

type Config struct {
	AppMode string `env:"APP_MODE" envDefault:"prod"`
}

func NewConfig() (*Config, error) {
	conf := &Config{}
	if err := env.Parse(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func (conf *Config) IsDevMode() bool {
	return conf.AppMode == "dev"
}
