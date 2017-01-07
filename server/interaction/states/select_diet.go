package states

import(
	. "github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
)
type SelectDiet struct{
	messengerProvider providers.IMessengerProvider
	userContext context.IUserContext
	stateFactory IStateFactory
}

const SELECT_DIET_STATE_ID = 10

func NewSelectDiet(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory IStateFactory) *SelectDiet {
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
	quickReplies[DIET_VEGAN_TITLE] = DIET_VEGAN_INPUT
	quickReplies[DIET_VEGETARIAN_TITLE] = DIET_VEGETARIAN_INPUT
	quickReplies[DIET_ANYTHING_TITLE] = DIET_ANYTHING_INPUT
	text := "אז.. מה אתה?"
	if !state.userContext.IsMale(){
		text = "אז.. מה את?"
	}

	return state.messengerProvider.SendQuickReplyMessage(externalUserID, text, quickReplies)
}

func (state *SelectDiet) Next(input IStateInput) (nextState IState, err error) {
	payload := input.Payload()
	var nextStateID int64
	switch payload {
	case DIET_ANYTHING_INPUT:
		nextStateID = SELECTED_ANYTHING_DIET_STATE_ID
	case DIET_VEGAN_INPUT:
		nextStateID = SELECTED_VEGAN_DIET_STATE_ID
	case DIET_VEGETARIAN_INPUT:
		nextStateID = SELECTED_VEGETARIAN_DIET_STATE_ID
	default:
		nextStateID = SELECT_DIET_OR_NOT_STATE_ID
	}

	return state.stateFactory.GetState(nextStateID)
}
