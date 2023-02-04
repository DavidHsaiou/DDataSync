package config

type Config struct {
	env Env `env:"APP_ENV" default:"dev"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) GetEnv() Env {
	return c.env
}
