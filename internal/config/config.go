package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port                string
	OpenRouterAPIKey    string
	OpenRouterBaseURL   string
	OpenRouterModel     string
	OpenRouterFallbacks string
	ConfidenceThreshold string
	RedisUsername       string
	RedisPassword       string
	RedisPort           string
	RedisURL            string
	PostgresDSN         string
}

func LoadConfig() (*Config, error) {
	redisUsername := getEnv("REDIS_USERNAME", "default")
	redisPassword := getEnv("REDIS_PASSWORD", "")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisURL := fmt.Sprintf("redis://%s:%s@%s:%s", redisUsername, redisPassword, redisHost, redisPort)

	pgHost := getEnv("POSTGRES_HOST", "localhost")
	pgUser := getEnv("POSTGRES_USER", "user")
	pgPassword := getEnv("POSTGRES_PASSWORD", "password")
	pgDB := getEnv("POSTGRES_DB", "otosei")
	pgPort := getEnv("POSTGRES_PORT", "5432")
	postgresDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		pgUser, pgPassword, pgHost, pgPort, pgDB,
	)

	config := &Config{
		Port:                getEnv("PORT", "8080"),
		OpenRouterAPIKey:    getEnv("OPENROUTER_API_KEY", ""),
		OpenRouterBaseURL:   getEnv("OPENROUTER_BASE_URL", "https://openrouter.ai/api/v1/chat/completions"),
		OpenRouterModel:     getEnv("OPENROUTER_MODEL", "deepseek/deepseek-r1:free"),
		OpenRouterFallbacks: getEnv("OPENROUTER_FALLBACKS", "deepseek/deepseek-chat-v3-0324:free,google/gemma-3n-e2b-it:free"),
		ConfidenceThreshold: getEnv("CONFIDENCE_THRESHOLD", "0.7"),
		RedisUsername:       redisUsername,
		RedisPassword:       redisPassword,
		RedisPort:           redisPort,
		RedisURL:            redisURL,
		PostgresDSN:         postgresDSN,
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
