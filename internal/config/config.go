package config

import (
	"os"
	"strconv"
)

type Config struct {
	App   AppConfig
	Mongo MongoConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type MongoConfig struct {
	URI            string
	Database       string
	TimeoutSeconds int
}

func Load() Config {
	return Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "go-api-mongo-kit"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Mongo: MongoConfig{
			URI:            getEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database:       getEnv("MONGO_DATABASE", "go_api_mongo_kit"),
			TimeoutSeconds: getEnvAsInt("MONGO_TIMEOUT_SECONDS", 5),
		},
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsedValue
}
