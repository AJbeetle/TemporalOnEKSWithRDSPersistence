package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TemporalHost string
	Port         string
}

func Load() *Config {
	// Loading .env file — only for local development
	// In Kubernetes, real env vars are injected and this is safely ignored
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system environment variables")
	}

	return &Config{
		TemporalHost: getEnv("TEMPORAL_HOST", "temporal-frontend.temporal.svc.cluster.local:7233"),
		Port:         getEnv("PORT", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
