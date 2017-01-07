package states

import(
	. "github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
)

type SelectSensitivityOrNot struct{
	messengerProvider providers.IMessengerProvider
	userContext context.IUserContext
	stateFactory IStateFactory
}

const SELECT_SENSITIVITY_OR_NOT_STATE_ID = 22

func NewSelectSensitivityOrNot(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory IStateFactory) *SelectSensitivityOrNot {
	return &SelectSensitivityOrNot{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectSensitivityOrNot) ID() int64 {
	return SELECT_SENSITIVITY_OR_NOT_STATE_ID
}

func (state *SelectSensitivityOrNot) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()

	quickReplies := make(map[string]string)
	quickReplies[NO_SENSITIVITY_TITLE] = NO_SENSITIVITY_INPUT
	quickReplies[GLUTEN_SENSITIVITY_TITLE] = GLUTEN_SENSITIVITY_INPUT
	quickReplies[MILK_SENSITIVITY_TITLE] = MILK_SENSITIVITY_INPUT

	text := "אם אתה רגיש למשהו חבל שנתפספס :) רגיש למשהו?"
	if !state.userContext.IsMale(){
		text = "אם את רגישה למשהו חבל שנתפספס :) רגישה למשהו?"
	}

	return state.messengerProvider.SendQuickReplyMessage(externalUserID, text, quickReplies)
}

func (state *SelectSensitivityOrNot) Next(input IStateInput) (nextState IState, err error) {
	payload := input.Payload()
	var nextStateID int64
	switch payload {
	case NO_SENSITIVITY_INPUT:
		nextStateID = SELECTED_NO_SENSITIVITY_STATE_ID
	case GLUTEN_SENSITIVITY_INPUT:
		nextStateID = SELECTED_GLUTEN_SENSITIVITY_STATE_ID
	case MILK_SENSITIVITY_INPUT:
		nextStateID = SELECTED_MILK_SENSITIVITY_STATE_ID
	case FREE_TEXT_INPUT:
		nextStateID = DID_NOT_SELECTED_SENSITIVITY_STATE_ID
	default:
		return nil, nil
	}

	return state.stateFactory.GetState(nextStateID)
}

func (state *SelectSensitivityOrNot) GetNextStage() (IState, error) {
	return nil, nil
}
