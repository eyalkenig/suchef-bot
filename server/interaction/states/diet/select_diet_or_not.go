package diet

import (
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs/diet"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
)

type SelectDietOrNot struct {
	userContext       context.IUserContext
	messengerProvider providers.IMessengerProvider
	stateFactory      interfaces.IStateFactory
}

const SELECT_DIET_OR_NOT_STATE_ID = 12

func NewSelectDietOrNot(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectDietOrNot {
	return &SelectDietOrNot{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectDietOrNot) ID() int64 {
	return SELECT_DIET_OR_NOT_STATE_ID
}

func (state *SelectDietOrNot) Act() (err error) {
	quickReplies := make(map[string]string)
	quickReplies[diet.DIET_VEGAN_TITLE] = diet.DIET_VEGAN_INPUT
	quickReplies[diet.DIET_VEGETARIAN_TITLE] = diet.DIET_VEGETARIAN_INPUT
	quickReplies[diet.DIET_ANYTHING_TITLE] = diet.DIET_ANYTHING_INPUT
	text := "אני עדיין קצת מתקשה בעברית, אני צריכה שתבחר מהאופציות :) אז.. איך תגדיר את עצמך?"
	if !state.userContext.IsMale() {
		text = "אני עדיין קצת מתקשה בעברית, אני צריכה שתבחרי מהאופציות :) אז.. איך תגדירי את עצמך?"
	}
	return state.messengerProvider.SendQuickReplyMessage(state.userContext.GetExternalUserID(), text, quickReplies)
}

func (state *SelectDietOrNot) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
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
		nextStateID = DID_NOT_SELECTED_DIET_STATE_ID
	default:
		return nil, nil
	}

	return state.stateFactory.GetState(nextStateID)
}

func (state *SelectDietOrNot) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
