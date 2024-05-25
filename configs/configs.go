package Configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Global struct {
	TelegramToken      string
	HostAddress        string
	EmailSufixTelegram string
}

var GlobalConfigs Global

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	GlobalConfigs.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	GlobalConfigs.HostAddress = os.Getenv("ST_SMTP_LISTEN")
	GlobalConfigs.EmailSufixTelegram = os.Getenv("EMAIL_SUFIX_TELEGRAM")
}
