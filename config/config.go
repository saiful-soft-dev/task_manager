package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func LoadConfig() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
