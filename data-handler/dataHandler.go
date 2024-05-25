package DataHandler

import (
	"io"
	"log"
	"strings"

	"github.com/mnako/letters"
)

type DataHandlerInterface interface {
	OnMailCreated(data []byte, from string, to []string)
	OnMailReceived(r io.Reader) error
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

// func (h *DataHandlerStruct) OnMailReceived(r io.Reader) {
// 	m, err := mail.ReadMessage(r) // .ReadMessage(os.Stdin)
// 	if err != nil {
// 		log.Fatalln("Parse mail KO -", err)
// 	}

// 	// Display only the main headers of the message. The "From","To" and "Subject" headers
// 	// have to be decoded if they were encoded using RFC 2047 to allow non ASCII characters.
// 	// We use a mime.WordDecode for that.
// 	dec := new(mime.WordDecoder)
// 	from, _ := dec.DecodeHeader(m.Header.Get("From"))
// 	to, _ := dec.DecodeHeader(m.Header.Get("To"))
// 	subject, _ := dec.DecodeHeader(m.Header.Get("Subject"))
// 	fmt.Println("From:", from)
// 	fmt.Println("To:", to)
// 	fmt.Println("Date:", m.Header.Get("Date"))
// 	fmt.Println("Subject:", subject)
// 	fmt.Println("Content-Type:", m.Header.Get("Content-Type"))
// 	fmt.Println()

// 	mediaType, params, err := mime.ParseMediaType(m.Header.Get("Content-Type"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if !strings.HasPrefix(mediaType, "multipart/") {
// 		log.Fatalf("Not a multipart MIME message\n")
// 	}

// 	// Recursivey parsed the MIME parts of the Body, starting with the first
// 	// level where the MIME parts are separated with params["boundary"].
// 	ParsePart(m.Body, params["boundary"], 1)
// }

func (h *DataHandlerStruct) OnMailReceived(r io.Reader) error {
	email, err := letters.ParseEmail(r)
	if err != nil {
		// log.Fatal(err)
		return err
	}
	log.Println("*************           KSA Mail To Telegram Begin           ********************")
	log.Println("***********       FROM    : ", email.Headers.From)
	log.Println("***********       TO      : ", email.Headers.Sender)
	log.Println("***********       SUBJECT : ", email.Headers.Subject)
	log.Println("***********       DATA    : ", email.Text)
	log.Println("*************           KSA Mail To Telegram END             ********************")

	return nil
}
