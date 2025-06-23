package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	MongoDB MongoDBConfig `yaml:"mongodb"`
	Server  ServerConfig
}

type MongoDBConfig struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
}

type ServerConfig struct {
	Port string
}

func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	return &Config{
		MongoDB: MongoDBConfig{
			URI:      getEnv("MONGO_URI", "mongodb://localhost:27017/homeassistant"),
			Database: getEnv("MONGO_DATABASE_NAME", "homeassistant"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
