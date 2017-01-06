package server

import (
	"fmt"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

type SuchefServer struct {
	messengerClient *messenger.Messenger
}

func NewSuchefServer(messengerClient *messenger.Messenger) *SuchefServer{
	return &SuchefServer{messengerClient: messengerClient}
}

func (suchefServer *SuchefServer) BindMessageReceived() messenger.MessageReceivedHandler{
	return func (event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		profile, err := suchefServer.messengerClient.GetProfile(opts.Sender.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		resp, err := suchefServer.messengerClient.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Hello, %s %s, %s", profile.FirstName, profile.LastName, msg.Text))

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", resp)
	}
}


