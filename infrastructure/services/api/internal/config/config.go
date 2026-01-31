package config

type Config struct {
}

func New(getenv func(string) string) (*Config, error) {
	return &Config{}, nil
}
