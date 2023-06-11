package config

type Config struct {
	DefaultOffset int
	DefaultLimit  int
}

func Load() Config {
	cfg := Config{}
	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10
	return cfg
}
