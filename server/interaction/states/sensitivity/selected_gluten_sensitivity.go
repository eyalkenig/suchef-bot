package sensitivity

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/theme"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
)

type SelectedGlutenSensitivity struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECTED_GLUTEN_SENSITIVITY_STATE_ID = 24
const GLUTEN_SENSITIVITY_TYPE_ID = 10

func NewSelectedGlutenSensitivity(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectedGlutenSensitivity {
	return &SelectedGlutenSensitivity{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedGlutenSensitivity) ID() int64 {
	return SELECTED_GLUTEN_SENSITIVITY_STATE_ID
}

func (state *SelectedGlutenSensitivity) Act() (err error) {
	err = state.userContext.SetSensitivity(GLUTEN_SENSITIVITY_TYPE_ID)
	if err != nil {
		return err
	}
	externalUserID := state.userContext.GetExternalUserID()
	message := "בסדר גמור! חשוב לי לציין: יש מנות ללא גלוטן, אבל המטבח עצמו עלול להכיל עקבות גלוטן.. מקווה שלא ביאסתי..!"
	return state.messengerProvider.SendSimpleMessage(externalUserID, message)
}

func (state *SelectedGlutenSensitivity) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedGlutenSensitivity) GetNextStage() (interfaces.IState, error) {
	return state.stateFactory.GetState(theme.SELECT_THEME_STATE_ID)
}
