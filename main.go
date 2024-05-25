package main

import (
	"log"

	DataHandler "ksa-smtp-telegram/data-handler"
	SmtpServer "ksa-smtp-telegram/smtp-server"
)

// var addr = "127.0.0.1:1025"
var addr = "0.0.0.0:1025"

func main() {

	dtHandler := DataHandler.MyDataHandler{}
	SmtpServer.SetDataMailHandler(&dtHandler)

	SmtpServer.SetConfig(addr, "localhost", true)

	log.Println("Starting SMTP server at", addr)
	log.Fatal(SmtpServer.ListenAndServe())

}
