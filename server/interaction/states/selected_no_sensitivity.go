package states

import(
	. "github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
)

type SelectedNoSensitivity struct{
	messengerProvider providers.IMessengerProvider
	userContext context.IUserContext
	stateFactory IStateFactory
}

const SELECTED_NO_SENSITIVITY_STATE_ID = 28
const NO_SENSITIVITY_TYPE_ID = 0

func NewSelectedNoSensitivity(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory IStateFactory) *SelectedNoSensitivity {
	return &SelectedNoSensitivity{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedNoSensitivity) ID() int64 {
	return SELECTED_NO_SENSITIVITY_STATE_ID
}

func (state *SelectedNoSensitivity) Act() (err error) {
	return state.userContext.SetSensitivity(NO_SENSITIVITY_TYPE_ID)
}

func (state *SelectedNoSensitivity) Next(input IStateInput) (nextState IState, err error) {
	return nil, nil
}

func (state *SelectedNoSensitivity) GetNextStage() (IState, error) {
	return nil, nil
}
