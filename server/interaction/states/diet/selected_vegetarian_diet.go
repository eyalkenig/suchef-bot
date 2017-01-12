package diet

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/sensitivity"
)

type SelectedVegetarianDiet struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECTED_VEGETARIAN_DIET_STATE_ID = 18
const VEGETARIAN_DIET_TYPE_ID = 20

func NewSelectedVegetarianDiet(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectedVegetarianDiet {
	return &SelectedVegetarianDiet{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedVegetarianDiet) ID() int64 {
	return SELECTED_VEGETARIAN_DIET_STATE_ID
}

func (state *SelectedVegetarianDiet) Act() (err error) {
	err = state.userContext.SetDiet(VEGETARIAN_DIET_TYPE_ID)
	if err != nil {
		return err
	}

	externalUserID := state.userContext.GetExternalUserID()
	message := "מעולה! יש אחלה תבשילים לצמחוניים 😋"
	return state.messengerProvider.SendSimpleMessage(externalUserID, message)
}

func (state *SelectedVegetarianDiet) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedVegetarianDiet) GetNextStage() (interfaces.IState, error) {
	return state.stateFactory.GetState(sensitivity.SELECT_SENSITIVITY_STATE_ID)
}
