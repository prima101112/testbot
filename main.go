package main

import (
	"github.com/prima101112/messengerbot"
	"log"
	"net/http"
	"os"
	"io"
	"fmt"
	neturl "net/url"
	"strings"
)

var Bot *messengerbot.MessengerBot

func init(){
	accessToken := "EAAQTZBKDffkkBADSQpcSJQncT5g5lPZBPb05BJhFag0q6BdkdTOblUwO612ec79ML5r6sUZCvxrAdZBohl0HZAAbD0reTBTfTVARtmYaesCbGkZBJ8gGLGZA2JjxTtxy8czYXljZCYccslmrmnelQsvH9XsuAJIw0yxKtRNpP1KCRAZDZD"
	verifyToken := "abcde"
	Bot = messengerbot.NewMessengerBot(accessToken, verifyToken)
	log.Print(Bot)
}

func main() {
	MReHand := messengerbot.MessageReceivedHandler(MRe)
	Bot.MessageReceived = MReHand
  Bot.Debug = true

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		Bot.Handler(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MRe(b *messengerbot.MessengerBot, event messengerbot.Event, mopts messengerbot.MessageOpts, recm messengerbot.ReceivedMessage) {
	log.Println("Attachment : ", recm.Attachments)
	log.Println("INCOMING : ", recm)
	user := messengerbot.NewUserFromId(mopts.Sender.ID)
	msg := messengerbot.NewMessage("Hola " + recm.Message.Text)

	if len(recm.Attachments) != 0 {
		if recm.Attachments[0].Type == "image" {
			atchpayload := recm.Attachments[0].Payload
			mapatch := atchpayload.(map[string]interface{})
			url := mapatch["url"].(string)
			DownloadImage(url)
		}
	}

  _, err := b.Send(user, msg, messengerbot.NotificationTypeRegular)
	log.Print(err)
}

func DownloadImage(url string){
    // don't worry about errors
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }

    defer response.Body.Close()

		//parse
		URL, err := neturl.Parse(url)
		if err != nil {
        log.Fatal(err)
    }
		arpath := strings.Split(URL.Path, "/")
		filename := arpath[len(arpath)-1]
    //open a file for writing
    file, err := os.Create("images/"+filename)
    if err != nil {
        log.Fatal(err)
    }
    // Use io.Copy to just dump the response body to the file. This supports huge files
    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
    fmt.Println("Download Success!")
}
