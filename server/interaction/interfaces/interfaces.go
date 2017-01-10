package interfaces

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

type IStateMachineController interface {
	InitUser() (err error)
	Handle(message messenger.ReceivedMessage) (err error)
}

type IState interface {
	Act() (err error)
	Next(input IStateInput) (nextState IState, err error)
	GetNextStage() (IState, error)
	ID() int64
}

type IStateDataProvider interface {
}

type IStateFactory interface {
	GetState(stateID int64) (state IState, err error)
	GetInitialState() (state IState)
}

type IStateInputFactory interface {
	CastMessageToInput(messenger.ReceivedMessage) (input IStateInput, err error)
}
type IStateInput interface {
	Payload() string
}
