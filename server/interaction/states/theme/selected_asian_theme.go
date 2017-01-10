package theme

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/providers"
)

type SelectedAsianTheme struct {
	messengerProvider providers.IMessengerProvider
	userContext       context.IUserContext
	stateFactory      interfaces.IStateFactory
}

const SELECTED_ASIAN_THEME_STATE_ID = 34
const ASIAN_THEME_TYPE_ID = 10

func NewSelectedAsianTheme(userContext context.IUserContext, messengerProvider providers.IMessengerProvider, stateFactory interfaces.IStateFactory) *SelectedAsianTheme {
	return &SelectedAsianTheme{userContext: userContext, messengerProvider: messengerProvider, stateFactory: stateFactory}
}

func (state *SelectedAsianTheme) ID() int64 {
	return SELECTED_ASIAN_THEME_STATE_ID
}

func (state *SelectedAsianTheme) Act() (err error) {
	externalUserID := state.userContext.GetExternalUserID()
	//err = state.messengerProvider.SendSimpleMessage(externalUserID, "דג בקארי")
	//if err != nil {
	//	return err
	//}
	//return state.messengerProvider.SendImage(externalUserID, "https://s28.postimg.org/u1hmefp1p/malai_not_grained.jpg")
	quickReplies := make(map[string]string)
	quickReplies["דג בקארי"] = "https://s28.postimg.org/u1hmefp1p/malai_not_grained.jpg"
	quickReplies["מאלאי כופתה"] = "https://s27.postimg.org/t7bx85c37/IMG_20160710_001110.jpg"
	quickReplies["באן במילוי משתנה"] = "https://s28.postimg.org/tvyqixb2l/IMG_20161225_225502.jpg"

	return state.messengerProvider.SendGenericTemplate(externalUserID, quickReplies)
}

func (state *SelectedAsianTheme) Next(input interfaces.IStateInput) (nextState interfaces.IState, err error) {
	return nil, nil
}

func (state *SelectedAsianTheme) GetNextStage() (interfaces.IState, error) {
	return nil, nil
}
