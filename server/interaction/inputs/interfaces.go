package inputs

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

type IStateInputFactory interface{
	CastMessageToInput(messenger.ReceivedMessage) (input IStateInput, err error)
}
type IStateInput interface{
	Payload() string
}
