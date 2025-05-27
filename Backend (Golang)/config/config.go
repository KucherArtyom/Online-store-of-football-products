package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Env       string
		Port      string
		JWTSecret string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	AppConfig.App.Env = GetEnv("APP_ENV", "development")
	AppConfig.App.Port = GetEnv("SERVER_PORT", "8080")
	AppConfig.App.JWTSecret = GetEnv("JWT_SECRET", "supersecretkey") // Добавляем секрет для JWT
	AppConfig.Database.Host = GetEnv("DB_HOST", "localhost")
	AppConfig.Database.Port = GetEnv("DB_PORT", "5432")
	AppConfig.Database.User = GetEnv("DB_USER", "postgres")
	AppConfig.Database.Password = GetEnv("DB_PASSWORD", "")
	AppConfig.Database.Name = GetEnv("DB_NAME", "footballstore")
}

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
