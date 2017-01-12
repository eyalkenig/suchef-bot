package diet

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/sensitivity"
)

type SelectedVeganDiet struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECTED_VEGAN_DIET_STATE_ID = 16
const VEGAN_DIET_TYPE_ID = 10

func NewSelectedVeganDiet(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectedVeganDiet {
	return &SelectedVeganDiet{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedVeganDiet) ID() int64 {
	return SELECTED_VEGAN_DIET_STATE_ID
}

func (state *SelectedVeganDiet) Act() (err error) {
	err = state.userContext.SetDiet(VEGAN_DIET_TYPE_ID)
	if err != nil {
		return err
	}
	externalUserID := state.userContext.GetExternalUserID()
	message := "מגניב! האמת שיש לי מחשבות להתקדם לטבעונות.. בנתיים אני לומדת לבשל עוד דברים טבעוניים 😋"
	return state.messengerProvider.SendSimpleMessage(externalUserID, message)
}

func (state *SelectedVeganDiet) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedVeganDiet) GetNextStage() (interfaces.IState, error) {
	return state.stateFactory.GetState(sensitivity.SELECT_SENSITIVITY_STATE_ID)
}
