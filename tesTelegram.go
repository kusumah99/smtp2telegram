package main

import (
	Configs "ksa-smtp-telegram/configs"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func TesTelegram() {
	log.Println(Configs.GlobalConfigs.TelegramToken)

	bot, err := tgbotapi.NewBotAPI(Configs.GlobalConfigs.TelegramToken)
	if err != nil {
		// log.Panic(err)
		log.Println(err)
		return
	}
	idChat, err := strconv.ParseInt("658662055", 10, 64)
	if err != nil {
		log.Println("error ", err)
		os.Exit(1)
	}
	// msg := tgbotapi.NewMessage(idChat, "Message from Notitication API:\n"+message)
	msg := tgbotapi.NewMessage(idChat, "Message from Notitication API:\n\nIni pesannya")
	msgbot, err := bot.Send(msg)
	if err != nil {
		log.Println("error2 ", err)
		os.Exit(1)
	}

	log.Println("Namanya: " + msgbot.From.UserName)

}
