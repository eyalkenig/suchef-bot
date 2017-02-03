package theme

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/repositories"
)

type SelectedMoroccanTheme struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
	courseRepository  repositories.ICourseRepository
}

const SELECTED_MOROCCAN_THEME_STATE_ID = 36
const MOROCCAN_THEME_TYPE_ID = 20

func NewSelectedMoroccanTheme(userContext context.IUserContext,
	messengerProvider providers.IMessengerProvider,
	stateFactory interfaces.IStateFactory,
	courseRepository repositories.ICourseRepository) *SelectedMoroccanTheme {
	return &SelectedMoroccanTheme{userContext: userContext,
		messengerProvider: messengerProvider,
		stateFactory:      stateFactory,
		courseRepository:  courseRepository}
}

func (state *SelectedMoroccanTheme) ID() int64 {
	return SELECTED_MOROCCAN_THEME_STATE_ID
}

func (state *SelectedMoroccanTheme) Act() error {
	externalUserID := state.userContext.GetExternalUserID()
	t := "moroccan"
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

func (state *SelectedMoroccanTheme) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedMoroccanTheme) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
