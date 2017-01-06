package main

import (
	"net/http"
	"github.com/eyalkenig/suchef-bot/server"

	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"os"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"fmt"
)

func main() {
	messenger := &messenger.Messenger{
		VerifyToken: os.Getenv("VERIFY_TOKEN"),
		AppSecret: os.Getenv("APP_SECRET"),
		AccessToken: os.Getenv("PAGE_ACCESS_KEY"),
	}

	dbConnectionParams := providers.DBConnectionParams{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Address: os.Getenv("DB_ADDRESS"),
		DBName: os.Getenv("DB_NAME"),
	}

	accountID := int64(1)

	suchefServer, err := server.NewSuchefServer(accountID, messenger, dbConnectionParams)

	if err != nil {
		fmt.Println("could not create suchef server. error: " + err.Error())
	}

	messenger.MessageReceived = suchefServer.BindMessageReceived()

	http.HandleFunc("/webhook", messenger.Handler)
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
