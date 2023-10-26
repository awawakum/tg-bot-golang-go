package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramConfig
	DBConfig
	APIConfig
}

type TelegramConfig struct {
	TelegramToken string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type APIConfig struct {
	ApiUrl string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	tg := TelegramConfig{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
	}

	db := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	api := APIConfig{
		ApiUrl: os.Getenv("API_URL"),
	}

	return &Config{
		DBConfig:       db,
		TelegramConfig: tg,
		APIConfig:      api,
	}, nil
}
