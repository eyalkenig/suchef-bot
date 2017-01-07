package interaction

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

type IStateMachineController interface {
	InitUser() (err error)
	Handle(message messenger.ReceivedMessage) (err error)
}
