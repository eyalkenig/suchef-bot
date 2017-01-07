package states

import(
	. "github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
)

type SelectedMoroccanTheme struct{
	messengerProvider providers.IMessengerProvider
	userContext context.IUserContext
	stateFactory IStateFactory
}

const SELECTED_MOROCCAN_THEME_STATE_ID = 36
const MOROCCAN_THEME_TYPE_ID = 20

func NewSelectedMoroccanTheme(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory IStateFactory) *SelectedMoroccanTheme {
	return &SelectedMoroccanTheme{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedMoroccanTheme) ID() int64 {
	return SELECTED_MOROCCAN_THEME_STATE_ID
}

func (state *SelectedMoroccanTheme) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()
	err = state.messengerProvider.SendSimpleMessage(externalUserID, "דג מרוקאי")
	if err != nil {
		return err
	}
	return state.messengerProvider.SendImage(externalUserID, "https://s30.postimg.org/49fugw4z5/fish.jpg")
}

func (state *SelectedMoroccanTheme) Next(input IStateInput) (nextState IState, err error) {
	return nil, nil
}

func (state *SelectedMoroccanTheme) GetNextStage() (IState, error) {
	return nil, nil
}
