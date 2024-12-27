package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    ServerPort string
    DbURL      string
}

func LoadConfig() Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using default values")
    }

    return Config{
        ServerPort: getEnv("SERVER_PORT", "8080"),
        DbURL:      getEnv("DB_URL", "postgres://user:password@localhost:5432/dbname"),
    }
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}
