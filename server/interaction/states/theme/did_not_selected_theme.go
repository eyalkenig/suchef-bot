package theme

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"github.com/eyalkenig/suchef-bot/server/repositories"
)

type DidNotSelectedTheme struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
	courseRepository  repositories.ICourseRepository
}

const DID_NOT_SELECTED_THEME_STATE_ID = 39

func NewDidNotSelectedTheme(userContext context.IUserContext,
	messengerProvider providers.IMessengerProvider,
	stateFactory interfaces.IStateFactory,
	courseRepository repositories.ICourseRepository) *DidNotSelectedTheme {
	return &DidNotSelectedTheme{userContext: userContext,
		messengerProvider: messengerProvider,
		stateFactory:      stateFactory,
		courseRepository:  courseRepository}
}

func (state *DidNotSelectedTheme) ID() int64 {
	return DID_NOT_SELECTED_THEME_STATE_ID
}

func (state *DidNotSelectedTheme) Act() (err error) {
	err = state.messengerProvider.SendSimpleMessage(state.userContext.GetExternalUserID(), "יאללה אני כבר אחליט לבד!")
	if err != nil {
		return err
	}
	externalUserID := state.userContext.GetExternalUserID()
	err = state.messengerProvider.SendSimpleMessage(externalUserID, "דג בקארי")
	if err != nil {
		return err
	}
	return state.messengerProvider.SendImage(externalUserID, "https://s28.postimg.org/u1hmefp1p/malai_not_grained.jpg")
}

func (state *DidNotSelectedTheme) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *DidNotSelectedTheme) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
