package internal

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
}

func GetConfig() (*Config, error) {
	currentWorkDirectory, _ := os.Getwd()
	pathToEnv, err := filepath.Abs(filepath.Join(currentWorkDirectory, ".env"))
	if err != nil {
		return nil, err
	}

	err = godotenv.Load(pathToEnv)
	if err != nil {
		return nil, err
	}

	return &Config{
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}, nil
}
