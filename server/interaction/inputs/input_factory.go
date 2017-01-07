package inputs

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"errors"
)

type StateInputFactory struct {}

func NewStateInputFactory() *StateInputFactory{
	return &StateInputFactory{}
}

func (inputFactory *StateInputFactory) CastMessageToInput(message messenger.ReceivedMessage) (input IStateInput, err error) {
	if message.Text == LETS_START_FROM_SCRATCH_INPUT {
		return NewLetsStartFromScratch(), nil
	}

	quickReply := message.QuickReply
	if quickReply == nil {
		return NewFreeTextInput(message.Text), nil
	}

	switch quickReply.Payload {
	case DIET_ANYTHING_INPUT:
		input = NewDietAnything()
	case DIET_VEGAN_INPUT:
		input = NewDietVegan()
	case DIET_VEGETARIAN_INPUT:
		input = NewDietVegetarian()
	case GLUTEN_SENSITIVITY_INPUT:
		input = NewSensitivityGluten()
	case MILK_SENSITIVITY_INPUT:
		input = NewSensitivityMilk()
	case NO_SENSITIVITY_INPUT:
		input = NewSensitivityNo()
	}

	if input == nil {
		return nil, errors.New("Invalid payload: " + message.QuickReply.Payload)
	}

	return input, nil
}
