package theme

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/repositories"
)

type SelectedMoroccasianTheme struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
	courseRepository  repositories.ICourseRepository
}

const SELECTED_MOROCCASIAN_THEME_STATE_ID = 38
const MOROCCASIAN_THEME_TYPE_ID = 30

func NewSelectedMoroccasianTheme(userContext context.IUserContext,
	messengerProvider providers.IMessengerProvider,
	stateFactory interfaces.IStateFactory,
	courseRepository repositories.ICourseRepository) *SelectedMoroccasianTheme {
	return &SelectedMoroccasianTheme{userContext: userContext,
		messengerProvider: messengerProvider,
		stateFactory:      stateFactory,
		courseRepository:  courseRepository}
}

func (state *SelectedMoroccasianTheme) ID() int64 {
	return SELECTED_MOROCCASIAN_THEME_STATE_ID
}

func (state *SelectedMoroccasianTheme) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()
	t := "moroccasian"
	theme, err := models.GetThemeByName(&t)
	if err != nil {
		return err
	}
	cards, err := GetSelectedThemeCourses(state.courseRepository, theme, state.userContext)

	if err != nil {
		return err
	}
	alternatives := GetAlterantivesCourses()
	return state.messengerProvider.SendGenericTemplate(externalUserID, alternatives, cards)
}

func (state *SelectedMoroccasianTheme) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedMoroccasianTheme) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
