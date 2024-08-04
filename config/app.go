package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DBType string
    DSN    string
    Port   string
}

var AppConfig Config

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    AppConfig = Config{
        DBType: getEnv("DB_TYPE", "postgres"),
        DSN:    getEnv("DB_DSN", ""),
        Port:   getEnv("PORT", "3000"),
    }
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}

