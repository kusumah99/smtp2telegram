package DataHandler

import (
	"log"
	"strings"
)

type DataHandlerInterface interface {
	OnMailCreated(data []byte, from string, to []string)
}
type DataHandlerStruct struct{}

func (h *DataHandlerStruct) OnMailCreated(data []byte, from string, to []string) {
	// Implementasi logika untuk memproses data

	log.Println("****************************************************************************")
	log.Println("****************************************************************************")
	log.Println("*************           KSA Mail To Telegram            ********************")
	log.Println("***********")
	log.Println("***********       FROM: ", from)
	log.Println("***********       TO:   ", strings.Join(to, ", "))
	log.Println("***********       DATA: ", string(data))
	log.Println("***********")
	log.Println("****************************************************************************")
	log.Println("****************************************************************************")

}
