package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port      string
    JWTSecret string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, reading environment variables directly")
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default port
    }

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET environment variable is required")
    }

    return &Config{
        Port:      port,
        JWTSecret: jwtSecret,
    }
}
