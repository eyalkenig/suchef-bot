package handlers

import (
	"fmt"
	"github.com/eyalkenig/suchef-bot/server/interaction/commands/inputs"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"strings"
)

type ShowIngredients struct {
	command           *inputs.ShowIngredients
	messengerProvider providers.IMessengerProvider
	botDataProvider   providers.IBotDataProvider
	userContext       context.IUserContext
}

func NewShowIngredients(command *inputs.ShowIngredients,
	messengerProvider providers.IMessengerProvider,
	botDataProvider providers.IBotDataProvider,
	userContext context.IUserContext) *ShowIngredients {
	return &ShowIngredients{command: command, messengerProvider: messengerProvider, botDataProvider: botDataProvider, userContext: userContext}
}

func (handler *ShowIngredients) Handle() error {
	courseID := handler.command.CourseID
	courseName, err := handler.botDataProvider.FetchCourseName(courseID)
	if err != nil {
		return err
	}
	ingredients, err := handler.botDataProvider.FetchIngredients(courseID)
	if err != nil {
		return err
	}

	var names []string
	for _, ingredient := range ingredients {
		names = append(names, ingredient.Name)
	}
	all := strings.Join(names, ", ")

	externalUserID := handler.userContext.GetExternalUserID()
	err = handler.messengerProvider.SendSimpleMessage(externalUserID, fmt.Sprintf("%s:", courseName))
	if err != nil {
		return err
	}
	return handler.messengerProvider.SendSimpleMessage(externalUserID, all)
}
