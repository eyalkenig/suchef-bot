package sensitivity

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/theme"
)

type SelectedMilkSensitivity struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECTED_MILK_SENSITIVITY_STATE_ID = 26
const MILK_SENSITIVITY_TYPE_ID = 20

func NewSelectedMilkSensitivity(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectedMilkSensitivity {
	return &SelectedMilkSensitivity{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedMilkSensitivity) ID() int64 {
	return SELECTED_MILK_SENSITIVITY_STATE_ID
}

func (state *SelectedMilkSensitivity) Act() (err error) {
	return state.userContext.SetSensitivity(MILK_SENSITIVITY_TYPE_ID)
}

func (state *SelectedMilkSensitivity) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedMilkSensitivity) GetNextStage() (interfaces.IState, error) {
	return state.stateFactory.GetState(theme.SELECT_THEME_STATE_ID)
}
