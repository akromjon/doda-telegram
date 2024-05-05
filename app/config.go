package app

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	ChatID   string
}

func Boot() (Config, error) {

	err := godotenv.Load("/usr/local/bin/.env")

	if err != nil {

		return Config{}, errors.New("cannot load .env file")

	}

	config := Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		ChatID:   os.Getenv("CHAT_ID"),
	}

	return config, nil

}
