package Configs

import (
	"os"
	// "github.com/joho/godotenv"
)

type Global struct {
	TelegramToken       string
	ListenAddress       string
	EmailDomainTelegram string
}

var GlobalConfigs Global

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// Gak jadi dari file .env langsung tapi via ENV host atau ENV docker/docker compose saja biar aman tokennya
	GlobalConfigs.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	GlobalConfigs.ListenAddress = os.Getenv("SMTP_LISTEN_ADDRESS")
	GlobalConfigs.EmailDomainTelegram = os.Getenv("EMAIL_DOMAIN_TELEGRAM")
	if len(GlobalConfigs.EmailDomainTelegram) < 5 {
		GlobalConfigs.EmailDomainTelegram = "kusumah99-tele.gram"
	}
}
