package theme

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/repositories"
)

type SelectedAsianTheme struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
	courseRepository  repositories.ICourseRepository
}

const SELECTED_ASIAN_THEME_STATE_ID = 34
const ASIAN_THEME_TYPE_ID = 10

func NewSelectedAsianTheme(userContext context.IUserContext,
	messengerProvider providers.IMessengerProvider,
	stateFactory interfaces.IStateFactory,
	courseRepository repositories.ICourseRepository) *SelectedAsianTheme {
	return &SelectedAsianTheme{userContext: userContext,
		messengerProvider: messengerProvider,
		stateFactory:      stateFactory,
		courseRepository:  courseRepository}
}

func (state *SelectedAsianTheme) ID() int64 {
	return SELECTED_ASIAN_THEME_STATE_ID
}

func (state *SelectedAsianTheme) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()
	t := "asian"
	theme, err := models.GetThemeByName(&t)
	if err != nil {
		return err
	}
	courses, err := GetSelectedThemeCourses(state.courseRepository, theme, state.userContext)

	if err != nil {
		return err
	}
	alternatives := GetAlterantivesCourses()
	return state.messengerProvider.SendGenericTemplate(externalUserID, alternatives, courses)
}

func (state *SelectedAsianTheme) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedAsianTheme) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
