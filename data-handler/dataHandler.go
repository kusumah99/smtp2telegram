package DataHandler

import (
	"log"
)

type MyDataHandler struct{}

func (h *MyDataHandler) OnMailCreated(data []byte, from string, to []string) {
	// Implementasi logika untuk memproses data

	log.Println("****************************************************************************")
	log.Println("****************************************************************************")
	log.Println("*************           KSA Mail To Telegram            ********************")
	log.Println("***********")
	log.Println("***********       FROM: ", from)
	log.Println("***********       TO:   ", to)
	log.Println("***********       DATA: ", data)
	log.Println("***********")
	log.Println("****************************************************************************")
	log.Println("****************************************************************************")

}
