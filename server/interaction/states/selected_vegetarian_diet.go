package states

import(
	. "github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
)

type SelectedVegetarianDiet struct{
	messengerProvider providers.IMessengerProvider
	userContext context.IUserContext
	stateFactory IStateFactory
}

const SELECTED_VEGETARIAN_DIET_STATE_ID = 18
const VEGETARIAN_DIET_TYPE_ID = 20

func NewSelectedVegetarianDiet(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory IStateFactory) *SelectedVegetarianDiet {
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

func (state *SelectedVegetarianDiet) Next(input IStateInput) (nextState IState, err error) {
	return nil, nil
}

func (state *SelectedVegetarianDiet) GetNextStage() (IState, error) {
	return state.stateFactory.GetState(SELECT_SENSITIVITY_STATE_ID)
}
