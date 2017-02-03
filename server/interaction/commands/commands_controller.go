package commands

import (
	"errors"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
)

type CommandsController struct {
	handlerFactory interfaces.ICommandsHandlerFactory
}

func NewCommandsController(messengerProvider providers.IMessengerProvider,
	botDataProvider providers.IBotDataProvider,
	userContext context.IUserContext) *CommandsController {
	handlerFactory := NewCommandsHandlerFactory(messengerProvider, botDataProvider, userContext)
	return &CommandsController{handlerFactory: handlerFactory}
}

func (controller *CommandsController) Handle(message messenger.ReceivedMessage) error {
	command, err := FetchCommand(message)
	if err != nil {
		return err
	}
	if command == nil {
		errors.New("unkown command. Text: " + message.Text)
	}

	handler, err := controller.handlerFactory.FetchHandler(command)

	if err != nil {
		return err
	}

	return handler.Handle()
}
