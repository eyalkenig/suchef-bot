package server

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"fmt"
)

type SuchefServer struct {
	controller ISuchefController
	accountID int64
}

func NewSuchefServer(accountID int64, messengerClient *messenger.Messenger, dbConnectionParams providers.DBConnectionParams) (server *SuchefServer, err error) {
	controller, err := NewSuchefController(messengerClient, dbConnectionParams)
	if err != nil {
		return nil, err
	}
	return &SuchefServer{controller: controller, accountID: accountID}, nil
}

func (suchefServer *SuchefServer) BindMessageReceived() messenger.MessageReceivedHandler{
	return func (event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		err := suchefServer.controller.Handle(suchefServer.accountID, event, opts, msg)
		if err != nil {
			fmt.Println("error handling message: " + err.Error())
		}
		fmt.Println("handled message: "+ msg.Text)
	}
}


