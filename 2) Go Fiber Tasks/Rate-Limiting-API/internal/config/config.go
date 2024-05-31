package config

import "os"

type Config struct {
	Server struct {
		Port string
	}
}

func Load() *Config {
	return &Config{
		Server: struct {
			Port string
		}{
			Port: os.Getenv("PORT"),
		},
	}
}
