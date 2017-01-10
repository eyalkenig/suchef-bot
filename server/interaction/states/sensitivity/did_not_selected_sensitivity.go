package sensitivity

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/theme"
)

type DidNotSelectedSensitivity struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const DID_NOT_SELECTED_SENSITIVITY_STATE_ID = 29

func NewDidNotSelectedSensitivity(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *DidNotSelectedSensitivity {
	return &DidNotSelectedSensitivity{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *DidNotSelectedSensitivity) ID() int64 {
	return DID_NOT_SELECTED_SENSITIVITY_STATE_ID
}

func (state *DidNotSelectedSensitivity) Act() (err error) {
	err = state.userContext.SetSensitivity(NO_SENSITIVITY_TYPE_ID)
	if err != nil {
		return err
	}
	return state.messengerProvider.SendSimpleMessage(state.userContext.GetExternalUserID(), "אוקיי לבנתיים אני אניח שאפשר הכל.. נמשיך!")
}

func (state *DidNotSelectedSensitivity) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *DidNotSelectedSensitivity) GetNextStage() (interfaces.IState, error) {
	return state.stateFactory.GetState(theme.SELECT_THEME_STATE_ID)
}
