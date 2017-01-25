package states

import (
	"errors"
	"fmt"

	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/diet"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/sensitivity"
	"github.com/eyalkenig/suchef-bot/server/interaction/states/theme"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/repositories"
)

type StateFactory struct {
	messengerProvider providers.IMessengerProvider
	dbProvider        providers.IBotDataProvider
	userContext       context.IUserContext
	courseRepository  repositories.ICourseRepository
}

func NewStateFactory(messengerProvider providers.IMessengerProvider, dbProvider providers.IBotDataProvider, userContext context.IUserContext, courseRepository repositories.ICourseRepository) *StateFactory {
	return &StateFactory{messengerProvider: messengerProvider, dbProvider: dbProvider, userContext: userContext, courseRepository: courseRepository}
}

func (stateFactory *StateFactory) GetState(stateID int64) (state interfaces.IState, err error) {
	switch stateID {
	case diet.SELECT_DIET_STATE_ID:
		return diet.NewSelectDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case diet.SELECT_DIET_OR_NOT_STATE_ID:
		return diet.NewSelectDietOrNot(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case diet.SELECTED_ANYTHING_DIET_STATE_ID:
		return diet.NewSelectedAnythingDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case diet.SELECTED_VEGAN_DIET_STATE_ID:
		return diet.NewSelectedVeganDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case diet.SELECTED_VEGETARIAN_DIET_STATE_ID:
		return diet.NewSelectedVegetarianDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case diet.DID_NOT_SELECTED_DIET_STATE_ID:
		return diet.NewDidNotSelectedDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case sensitivity.SELECT_SENSITIVITY_STATE_ID:
		return sensitivity.NewSelectSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case sensitivity.SELECT_SENSITIVITY_OR_NOT_STATE_ID:
		return sensitivity.NewSelectSensitivityOrNot(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case sensitivity.SELECTED_GLUTEN_SENSITIVITY_STATE_ID:
		return sensitivity.NewSelectedGlutenSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case sensitivity.SELECTED_MILK_SENSITIVITY_STATE_ID:
		return sensitivity.NewSelectedMilkSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case sensitivity.SELECTED_NO_SENSITIVITY_STATE_ID:
		return sensitivity.NewSelectedNoSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case sensitivity.DID_NOT_SELECTED_SENSITIVITY_STATE_ID:
		return sensitivity.NewDidNotSelectedSensitivity(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case theme.SELECT_THEME_STATE_ID:
		return theme.NewSelectFoodTheme(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case theme.SELECT_THEME_OR_NOT_STATE_ID:
		return theme.NewSelectFoodThemeOrNot(stateFactory.userContext, stateFactory.messengerProvider, stateFactory), nil
	case theme.SELECTED_ASIAN_THEME_STATE_ID:
		return theme.NewSelectedAsianTheme(stateFactory.userContext, stateFactory.messengerProvider, stateFactory, stateFactory.courseRepository), nil
	case theme.SELECTED_MOROCCAN_THEME_STATE_ID:
		return theme.NewSelectedMoroccanTheme(stateFactory.userContext, stateFactory.messengerProvider, stateFactory, stateFactory.courseRepository), nil
	case theme.SELECTED_MOROCCASIAN_THEME_STATE_ID:
		return theme.NewSelectedMoroccasianTheme(stateFactory.userContext, stateFactory.messengerProvider, stateFactory, stateFactory.courseRepository), nil
	case theme.DID_NOT_SELECTED_THEME_STATE_ID:
		return theme.NewDidNotSelectedTheme(stateFactory.userContext, stateFactory.messengerProvider, stateFactory, stateFactory.courseRepository), nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid state id: %d", stateID))
}

func (stateFactory *StateFactory) GetInitialState() (state interfaces.IState) {
	return diet.NewSelectDiet(stateFactory.userContext, stateFactory.messengerProvider, stateFactory)
}
