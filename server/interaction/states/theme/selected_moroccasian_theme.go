package theme

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
)

type SelectedMoroccasianTheme struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECTED_MOROCCASIAN_THEME_STATE_ID = 38
const MOROCCASIAN_THEME_TYPE_ID = 30

func NewSelectedMoroccasianTheme(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectedMoroccasianTheme {
	return &SelectedMoroccasianTheme{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedMoroccasianTheme) ID() int64 {
	return SELECTED_MOROCCASIAN_THEME_STATE_ID
}

func (state *SelectedMoroccasianTheme) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()
	err = state.messengerProvider.SendSimpleMessage(externalUserID, "קובה בקרם קוקוס וקארי")
	if err != nil {
		return err
	}
	return state.messengerProvider.SendImage(externalUserID, "https://s23.postimg.org/ccl4mikfv/kuba1.jpg")
}

func (state *SelectedMoroccasianTheme) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedMoroccasianTheme) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
