package interfaces

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

type IStateMachineController interface {
	InitUser() error
	Handle(message messenger.ReceivedMessage) error
}

type ICommandsController interface {
	Handle(message messenger.ReceivedMessage) error
}

type IInteractionController interface {
	InitUser() error
	Handle(message messenger.ReceivedMessage) error
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

type ICommand interface{}

type ICommandHandler interface {
	Handle() error
}
type ICommandsHandlerFactory interface {
	FetchHandler(command ICommand) (ICommandHandler, error)
}
