package sensitivity

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/theme"
)

type SelectedNoSensitivity struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECTED_NO_SENSITIVITY_STATE_ID = 28
const NO_SENSITIVITY_TYPE_ID = 0

func NewSelectedNoSensitivity(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectedNoSensitivity {
	return &SelectedNoSensitivity{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedNoSensitivity) ID() int64 {
	return SELECTED_NO_SENSITIVITY_STATE_ID
}

func (state *SelectedNoSensitivity) Act() (err error) {
	return state.userContext.SetSensitivity(NO_SENSITIVITY_TYPE_ID)
}

func (state *SelectedNoSensitivity) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedNoSensitivity) GetNextStage() (interfaces.IState, error) {
	return state.stateFactory.GetState(theme.SELECT_THEME_STATE_ID)
}
