package theme

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs/theme"
)

type SelectFoodTheme struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECT_THEME_STATE_ID = 30

func NewSelectFoodTheme(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectFoodTheme {
	return &SelectFoodTheme{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectFoodTheme) ID() int64 {
	return SELECT_THEME_STATE_ID
}

func (state *SelectFoodTheme) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()
	text := "שגעון! אני אזכור את ההעדפות האלו כדי לא לשגע אותך שוב.."
	err = state.messengerProvider.SendSimpleMessage(externalUserID, text)

	if err != nil {
		return err
	}

	quickReplies := make(map[string]string)
	quickReplies[theme.THEME_ASIAN_TITLE] = theme.THEME_ASIAN_INPUT
	quickReplies[theme.THEME_MOROCCAN_TITLE] = theme.THEME_MOROCCAN_INPUT
	quickReplies[theme.THEME_MOROCCASIAN_TITLE] = theme.THEME_MOROCCASIAN_INPUT
	text = "איזה סגנון אוכל מתחשק לך כעת?"

	return state.messengerProvider.SendQuickReplyMessage(externalUserID, text, quickReplies)
}

func (state *SelectFoodTheme) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	payload := input.Payload()
	var nextStateID int64
	switch payload {
	case theme.THEME_ASIAN_INPUT:
		nextStateID = SELECTED_ASIAN_THEME_STATE_ID
	case theme.THEME_MOROCCAN_INPUT:
		nextStateID = SELECTED_MOROCCAN_THEME_STATE_ID
	case theme.THEME_MOROCCASIAN_INPUT:
		nextStateID = SELECTED_MOROCCASIAN_THEME_STATE_ID
	case inputs.FREE_TEXT_INPUT:
		nextStateID = SELECT_THEME_OR_NOT_STATE_ID
	default:
		return nil, nil
	}

	return state.stateFactory.GetState(nextStateID)
}

func (state *SelectFoodTheme) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
