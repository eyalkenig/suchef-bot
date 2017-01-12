package diet

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/sensitivity"
)

type SelectedAnythingDiet struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECTED_ANYTHING_DIET_STATE_ID = 14
const ANYTHING_DIET_TYPE_ID = 0

func NewSelectedAnythingDiet(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectedAnythingDiet {
	return &SelectedAnythingDiet{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedAnythingDiet) ID() int64 {
	return SELECTED_ANYTHING_DIET_STATE_ID
}

func (state *SelectedAnythingDiet) Act() (err error) {
	return state.userContext.SetDiet(ANYTHING_DIET_TYPE_ID)
}

func (state *SelectedAnythingDiet) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedAnythingDiet) GetNextStage() (interfaces.IState, error) {
	return state.stateFactory.GetState(sensitivity.SELECT_SENSITIVITY_STATE_ID)
}
