package main

import (
	"github.com/eyalkenig/suchef-bot/server"
	"net/http"

	"fmt"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"os"
)

func main() {
	messenger := &messenger.Messenger{
		VerifyToken: os.Getenv("VERIFY_TOKEN"),
		AppSecret:   os.Getenv("APP_SECRET"),
		AccessToken: os.Getenv("PAGE_ACCESS_KEY"),
	}

	dbConnectionParams := providers.DBConnectionParams{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Address:  os.Getenv("DB_ADDRESS"),
		DBName:   os.Getenv("DB_NAME"),
	}

	accountID := int64(1)

	suchefServer, err := server.NewSuchefServer(accountID, messenger, dbConnectionParams)

	if err != nil {
		fmt.Println("could not create suchef server. error: " + err.Error())
	} else {
		fmt.Println("server started successfully")
	}

	messenger.MessageReceived = suchefServer.BindMessageReceived()
	messenger.Postback = suchefServer.BindPostbackReceived()

	http.HandleFunc("/webhook", messenger.Handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
