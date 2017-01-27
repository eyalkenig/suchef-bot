package sensitivity

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs/sensitivity"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
)

type SelectSensitivity struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECT_SENSITIVITY_STATE_ID = 20

func NewSelectSensitivity(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectSensitivity {
	return &SelectSensitivity{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectSensitivity) ID() int64 {
	return SELECT_SENSITIVITY_STATE_ID
}

func (state *SelectSensitivity) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()

	quickReplies := make(map[string]string)
	quickReplies[sensitivity.NO_SENSITIVITY_TITLE] = sensitivity.NO_SENSITIVITY_INPUT
	quickReplies[sensitivity.GLUTEN_SENSITIVITY_TITLE] = sensitivity.GLUTEN_SENSITIVITY_INPUT
	quickReplies[sensitivity.MILK_SENSITIVITY_TITLE] = sensitivity.MILK_SENSITIVITY_INPUT
	text := "אתה רגיש למשהו?"
	if !state.userContext.IsMale() {
		text = "את רגישה למשהו?"
	}

	return state.messengerProvider.SendQuickReplyMessage(externalUserID, text, quickReplies)
}

func (state *SelectSensitivity) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	payload := input.Payload()
	var nextStateID int64
	switch payload {
	case sensitivity.NO_SENSITIVITY_INPUT:
		nextStateID = SELECTED_NO_SENSITIVITY_STATE_ID
	case sensitivity.GLUTEN_SENSITIVITY_INPUT:
		nextStateID = SELECTED_GLUTEN_SENSITIVITY_STATE_ID
	case sensitivity.MILK_SENSITIVITY_INPUT:
		nextStateID = SELECTED_MILK_SENSITIVITY_STATE_ID
	case inputs.FREE_TEXT_INPUT:
		nextStateID = SELECT_SENSITIVITY_OR_NOT_STATE_ID
	default:
		return nil, nil
	}

	return state.stateFactory.GetState(nextStateID)
}

func (state *SelectSensitivity) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
