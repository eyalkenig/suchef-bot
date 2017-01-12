package inputs

import (
	"errors"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs/diet"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs/sensitivity"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs/theme"
)

type StateInputFactory struct{}

func NewStateInputFactory() *StateInputFactory {
	return &StateInputFactory{}
}

func (inputFactory *StateInputFactory) CastMessageToInput(message messenger.ReceivedMessage) (input interfaces.IStateInput, err error) {
	if message.Text == LETS_START_FROM_SCRATCH_INPUT {
		return NewLetsStartFromScratch(), nil
	}

	quickReply := message.QuickReply
	if quickReply == nil {
		return NewFreeTextInput(message.Text), nil
	}

	switch quickReply.Payload {
	case diet.DIET_ANYTHING_INPUT:
		input = diet.NewDietAnything()
	case diet.DIET_VEGAN_INPUT:
		input = diet.NewDietVegan()
	case diet.DIET_VEGETARIAN_INPUT:
		input = diet.NewDietVegetarian()
	case sensitivity.GLUTEN_SENSITIVITY_INPUT:
		input = sensitivity.NewSensitivityGluten()
	case sensitivity.MILK_SENSITIVITY_INPUT:
		input = sensitivity.NewSensitivityMilk()
	case sensitivity.NO_SENSITIVITY_INPUT:
		input = sensitivity.NewSensitivityNo()
	case theme.THEME_ASIAN_INPUT:
		input = theme.NewThemeAsian()
	case theme.THEME_MOROCCAN_INPUT:
		input = theme.NewThemeMoroccan()
	case theme.THEME_MOROCCASIAN_INPUT:
		input = theme.NewThemeMoroccasian()
	}

	if input == nil {
		return nil, errors.New("Invalid payload: " + message.QuickReply.Payload)
	}

	return input, nil
}
