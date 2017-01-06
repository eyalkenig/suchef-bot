package main

import (
	"net/http"
	"github.com/eyalkenig/suchef-bot/server"

	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

func main() {
	messenger := &messenger.Messenger{
		VerifyToken: "MY_VERIFY_TOKEN_SUCHEF",
		AppSecret: "2a551675367ec29d3cb5e043aa6f546c",
		AccessToken: "EAAVpMZASa6N8BAC3pH0I8oNLvNDGtZAgUT5CTdj37dSDhQgZAfhJGcZBQPiILqYzXZAWPQbpMkbiQmZBlfqPasmIWqQEW8qQz6zHTnmtuDBOac7QqDkqhpvZAkDnXgQkF1ZAa83TdyaOoduZAdFZBEg9RBJWMwhgFECBUL9aiInNTFQQZDZD",
	}

	suchefServer := server.NewSuchefServer(messenger)
	messenger.MessageReceived = suchefServer.BindMessageReceived()

	http.HandleFunc("/webhook", messenger.Handler)
	http.ListenAndServe(":3987", nil)
}
