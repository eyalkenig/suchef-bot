package diet

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs/diet"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
)

type SelectDiet struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECT_DIET_STATE_ID = 10

func NewSelectDiet(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectDiet {
	return &SelectDiet{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectDiet) ID() int64 {
	return SELECT_DIET_STATE_ID
}

func (state *SelectDiet) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()
	err = state.messengerProvider.SendSimpleMessage(externalUserID, "היי! קוראים לי סושף ואני כאן לעזור לך לגלות מה כדאי לך להזמין :)")
	if err != nil {
		return err
	}

	quickReplies := make(map[string]string)
	quickReplies[diet.DIET_VEGAN_TITLE] = diet.DIET_VEGAN_INPUT
	quickReplies[diet.DIET_VEGETARIAN_TITLE] = diet.DIET_VEGETARIAN_INPUT
	quickReplies[diet.DIET_ANYTHING_TITLE] = diet.DIET_ANYTHING_INPUT
	text := "אז.. מה אתה?"
	if !state.userContext.IsMale() {
		text = "אז.. מה את?"
	}

	return state.messengerProvider.SendQuickReplyMessage(externalUserID, text, quickReplies)
}

func (state *SelectDiet) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	payload := input.Payload()
	var nextStateID int64
	switch payload {
	case diet.DIET_ANYTHING_INPUT:
		nextStateID = SELECTED_ANYTHING_DIET_STATE_ID
	case diet.DIET_VEGAN_INPUT:
		nextStateID = SELECTED_VEGAN_DIET_STATE_ID
	case diet.DIET_VEGETARIAN_INPUT:
		nextStateID = SELECTED_VEGETARIAN_DIET_STATE_ID
	case inputs.FREE_TEXT_INPUT:
		nextStateID = SELECT_DIET_OR_NOT_STATE_ID
	default:
		return nil, nil
	}

	return state.stateFactory.GetState(nextStateID)
}

func (state *SelectDiet) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
