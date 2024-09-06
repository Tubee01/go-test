package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_LISTEN   string
	DATABASE_URL string
}

var AppConfig Config

// Config func to get env value
func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	AppConfig = Config{
		APP_LISTEN:   GetEnv("APP_LISTEN", ":3000"),
		DATABASE_URL: GetEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5433/postgres"),
	}
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
