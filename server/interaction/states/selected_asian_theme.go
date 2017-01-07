package states

import(
	. "github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
)

type SelectedAsianTheme struct{
	messengerProvider providers.IMessengerProvider
	userContext context.IUserContext
	stateFactory IStateFactory
}

const SELECTED_ASIAN_THEME_STATE_ID = 34
const ASIAN_THEME_TYPE_ID = 10

func NewSelectedAsianTheme(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory IStateFactory) *SelectedAsianTheme {
	return &SelectedAsianTheme{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedAsianTheme) ID() int64 {
	return SELECTED_ASIAN_THEME_STATE_ID
}

func (state *SelectedAsianTheme) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()
	err = state.messengerProvider.SendSimpleMessage(externalUserID, "דג בקארי")
	if err != nil {
		return err
	}
	return state.messengerProvider.SendImage(externalUserID, "https://s28.postimg.org/u1hmefp1p/malai_not_grained.jpg")
}

func (state *SelectedAsianTheme) Next(input IStateInput) (nextState IState, err error) {
	return nil, nil
}

func (state *SelectedAsianTheme) GetNextStage() (IState, error) {
	return nil, nil
}
