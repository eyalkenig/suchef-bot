package diet

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/sensitivity"
)

type DidNotSelectedDiet struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const DID_NOT_SELECTED_DIET_STATE_ID = 19

func NewDidNotSelectedDiet(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *DidNotSelectedDiet {
	return &DidNotSelectedDiet{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *DidNotSelectedDiet) ID() int64 {
	return SELECTED_ANYTHING_DIET_STATE_ID
}

func (state *DidNotSelectedDiet) Act() (err error) {
	err = state.userContext.SetDiet(ANYTHING_DIET_TYPE_ID)
	if err != nil {
		return err
	}
	return state.messengerProvider.SendSimpleMessage(state.userContext.GetExternalUserID(), "אוקיי לבנתיים אני אניח שאפשר הכל.. נמשיך!")
}

func (state *DidNotSelectedDiet) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *DidNotSelectedDiet) GetNextStage() (interfaces.IState, error) {
	return state.stateFactory.GetState(sensitivity.SELECT_SENSITIVITY_STATE_ID)
}
