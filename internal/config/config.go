package config

import (
	"os"
)

type Config struct {
	TemporalHost string
	Port         string
}

func Load() *Config {
	return &Config{
		TemporalHost: getEnv("TEMPORAL_HOST", "localhost:7233"),
		Port:         getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
