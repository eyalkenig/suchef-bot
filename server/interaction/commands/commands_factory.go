package commands

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"strings"
	"github.com/eyalkenig/suchef-bot/server/interaction/buttons"
	"github.com/eyalkenig/suchef-bot/server/interaction/commands/inputs"
	"strconv"
)

type CommandsFactory struct {}

func NewCommandsFactory() *CommandsFactory{
	return &CommandsFactory{}
}

func IsCommand(message messenger.ReceivedMessage) bool {
	if strings.HasPrefix(message.Text, buttons.INGREDIENTS_PAYLOAD) {
		return true
	}
	return false
}

func FetchCommand(message messenger.ReceivedMessage) (interfaces.ICommand, error) {
	if !IsCommand(message) {
		return nil, nil
	}

	if strings.HasPrefix(message.Text, buttons.INGREDIENTS_PAYLOAD) {
		courseIDstr := strings.TrimPrefix(message.Text, buttons.INGREDIENTS_PAYLOAD)
		courseID, err := strconv.Atoi(courseIDstr)
		if err != nil {
			return nil, err
		}
		return inputs.NewShowIngredients(int64(courseID)), nil
	}

	return nil, nil
}