package states

import (
	"errors"
	"fmt"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
)

type StateFactory struct{
	messengerProvider providers.IMessengerProvider
	dbProvider providers.IBotDataProvider
	userContext context.IUserContext
}

func NewStateFactory(messengerProvider providers.IMessengerProvider, dbProvider providers.IBotDataProvider, userContext context.IUserContext) *StateFactory{
	return &StateFactory{messengerProvider: messengerProvider, dbProvider: dbProvider, userContext: userContext}
}

func (stateFactory *StateFactory) GetState(stateID int64) (state IState, err error) {
	switch stateID {
	case SELECT_DIET_STATE_ID:
		return NewSelectDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECT_DIET_OR_NOT_STATE_ID:
		return NewSelectDietOrNot(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECTED_ANYTHING_DIET_STATE_ID:
		return NewSelectedAnythingDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECTED_VEGAN_DIET_STATE_ID:
		return NewSelectedVeganDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECTED_VEGETARIAN_DIET_STATE_ID:
		return NewSelectedVegetarianDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case DID_NOT_SELECTED_DIET_STATE_ID:
		return NewDidNotSelectedDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECT_SENSITIVITY_STATE_ID:
		return NewSelectSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECT_SENSITIVITY_OR_NOT_STATE_ID:
		return NewSelectSensitivityOrNot(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECTED_GLUTEN_SENSITIVITY_STATE_ID:
		return NewSelectedGlutenSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECTED_MILK_SENSITIVITY_STATE_ID:
		return NewSelectedMilkSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case SELECTED_NO_SENSITIVITY_STATE_ID:
		return NewSelectedNoSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case DID_NOT_SELECTED_SENSITIVITY_STATE_ID:
		return NewDidNotSelectedSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid state id: %d", stateID))
}

func (stateFactory *StateFactory) GetInitialState() (state IState) {
	return NewSelectDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory)
}