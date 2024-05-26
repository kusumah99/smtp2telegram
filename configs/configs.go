package Configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Global struct {
	TelegramToken       string
	ListenAddress       string
	EmailDomainTelegram string
}

var GlobalConfigs Global

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	GlobalConfigs.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	GlobalConfigs.ListenAddress = os.Getenv("SMTP_LISTEN_ADDRESS")
	GlobalConfigs.EmailDomainTelegram = os.Getenv("EMAIL_DOMAIN_TELEGRAM")
}
