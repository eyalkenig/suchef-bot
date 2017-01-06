package main

import (
	"net/http"
	"github.com/eyalkenig/suchef-bot/server"

	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"os"
)

func main() {
	messenger := &messenger.Messenger{
		VerifyToken: os.Getenv("VERIFY_TOKEN"),
		AppSecret: os.Getenv("APP_SECRET"),
		AccessToken: os.Getenv("PAGE_ACCESS_KEY"),
	}

	suchefServer := server.NewSuchefServer(messenger)
	messenger.MessageReceived = suchefServer.BindMessageReceived()

	http.HandleFunc("/webhook", messenger.Handler)
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
