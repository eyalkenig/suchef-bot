package commands

import (
	"errors"
	"fmt"
	"github.com/eyalkenig/suchef-bot/server/interaction/commands/handlers"
	"github.com/eyalkenig/suchef-bot/server/interaction/commands/inputs"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
)

type CommandsHandlerFactory struct {
	messengerProvider providers.IMessengerProvider
	botDataProvider   providers.IBotDataProvider
	userContext       context.IUserContext
}

func NewCommandsHandlerFactory(messengerProvider providers.IMessengerProvider,
	botDataProvider providers.IBotDataProvider,
	userContext context.IUserContext) *CommandsHandlerFactory {
	return &CommandsHandlerFactory{messengerProvider: messengerProvider, botDataProvider: botDataProvider, userContext: userContext}
}

func (factory *CommandsHandlerFactory) FetchHandler(command interfaces.ICommand) (interfaces.ICommandHandler, error) {
	switch castedCommand := command.(type) {
	case *inputs.ShowIngredients:
		return handlers.NewShowIngredients(castedCommand, factory.messengerProvider, factory.botDataProvider, factory.userContext), nil
	}

	return nil, errors.New(fmt.Sprintf("unkown command type: \t", command))
}
