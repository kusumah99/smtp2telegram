package main

import (
	Configs "ksa-smtp-telegram/configs"
	DataHandler "ksa-smtp-telegram/data-handler"
	SmtpServer "ksa-smtp-telegram/smtp-server"
	"log"
)

func main() {

	dtHandler := DataHandler.DataHandlerStruct{}
	SmtpServer.SetDataMailHandler(&dtHandler)

	addr := Configs.GlobalConfigs.ListenAddress

	// SmtpServer.SetConfig(addr, os.Stdout, true)
	SmtpServer.SetConfig(addr, nil, true)

	log.Println("Starting SMTP server at", addr)
	log.Fatal(SmtpServer.ListenAndServe())

}
