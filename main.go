package main

import (
	Configs "ksa-smtp-telegram/configs"
	DataHandler "ksa-smtp-telegram/data-handler"
	SmtpServer "ksa-smtp-telegram/smtp-server"
	"log"
)

func main() {

	log.Println()
	log.Println()
	log.Println("    ********************    SMTP To Telegram    ********************     ")
	log.Println("  **************************    Using Go    *************************    ")
	log.Println("By Kusumah Sasmita")
	log.Println()
	log.Println("ENV:")
	// 6859XXXXXX:XXXXXXXXXXX-XXXXXXXXXX-_XXXXXXDwAur
	c := Configs.GlobalConfigs.TelegramToken
	if len(c) > 30 {
		log.Println("TELEGRAM_TOKEN=", c[:4]+"XXXXXX:XXXXXXXXXXX-XXXXXXXXXX-_XXXXXX"+c[len(c)-5:])
	} else {
		log.Fatalln("INVALID TELEGRAM TOKEN")
	}
	log.Println("SMTP_LISTEN_ADDRESS=", Configs.GlobalConfigs.ListenAddress)
	log.Println("EMAIL_DOMAIN_TELEGRAM=", Configs.GlobalConfigs.EmailDomainTelegram)
	log.Println()

	dtHandler := DataHandler.DataHandlerStruct{}
	SmtpServer.SetDataMailHandler(&dtHandler)

	addr := Configs.GlobalConfigs.ListenAddress

	// SmtpServer.SetConfig(addr, os.Stdout, true)
	SmtpServer.SetConfig(addr, nil, true)

	log.Println("Starting SMTP server at", addr)
	log.Println()
	log.Fatal(SmtpServer.ListenAndServe())

}
