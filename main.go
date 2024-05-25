package main

import (
	"log"

	Configs "ksa-smtp-telegram/configs"
	DataHandler "ksa-smtp-telegram/data-handler"
	SmtpServer "ksa-smtp-telegram/smtp-server"
)

// var addr = "127.0.0.1:1025"
// var addr = "0.0.0.0:1025"

func main() {

	addr := Configs.GlobalConfigs.HostAddress
	// addr = os.Getenv("ST_SMTP_LISTEN")
	dtHandler := DataHandler.DataHandlerStruct{}
	SmtpServer.SetDataMailHandler(&dtHandler)

	SmtpServer.SetConfig(addr, "localhost", true)

	log.Println("Starting SMTP server at", addr)
	log.Fatal(SmtpServer.ListenAndServe())

}
