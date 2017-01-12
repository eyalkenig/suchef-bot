package server

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

type ISuchefController interface {
	Handle(accountID int64, event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) error
}
