package server

import (
	"fmt"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

type SuchefServer struct {
	controller ISuchefController
	accountID  int64
}

func NewSuchefServer(accountID int64, messengerClient *messenger.Messenger, dbConnectionParams providers.DBConnectionParams) (server *SuchefServer, err error) {
	controller, err := NewSuchefController(messengerClient, dbConnectionParams)
	if err != nil {
		return nil, err
	}
	return &SuchefServer{controller: controller, accountID: accountID}, nil
}

func (suchefServer *SuchefServer) BindMessageReceived() messenger.MessageReceivedHandler {
	return func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		err := suchefServer.controller.Handle(suchefServer.accountID, event, opts, msg)
		if err != nil {
			fmt.Println("error handling message: " + err.Error())
		}
		fmt.Println("handled message: " + msg.Text)
	}
}

func (suchefServer *SuchefServer) BindPostbackReceived() messenger.PostbackHandler {
	return func(event messenger.Event, opts messenger.MessageOpts, postback messenger.Postback) {
		fakeMid := fmt.Sprintf("postback_%d", event.ID)
		postbackMessage := messenger.ReceivedMessage{ID: fakeMid, Text: postback.Payload, Seq: -1}
		err := suchefServer.controller.Handle(suchefServer.accountID, event, opts, postbackMessage)
		if err != nil {
			fmt.Println("error handling postback: " + err.Error())
		}
		fmt.Println("handled posback: " + postback.Payload)
	}
}
