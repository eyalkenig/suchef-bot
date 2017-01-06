package server

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"fmt"
	"github.com/eyalkenig/suchef-bot/server/models"
	"errors"
)

type SuchefController struct {
	dataProvider providers.IBotDataProvider
	messengerClient *messenger.Messenger
}

func NewSuchefController(messengerClient *messenger.Messenger, dbConnectionParams providers.DBConnectionParams) (controller *SuchefController, err error) {
	dataProvider, err := providers.NewBotDataProvider(dbConnectionParams)
	if err != nil {
		return nil, err
	}
	return &SuchefController{dataProvider: dataProvider, messengerClient: messengerClient}, nil
}

func (controller *SuchefController) Handle(accountID int64, event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) (error){
	externalUserID := opts.Sender.ID
	user, err := controller.fetchUser(accountID, externalUserID)
	if (err != nil) {
		return err
	}

	if (user == nil) {
		return errors.New(fmt.Sprintf("user was not found. account id: %d, external user id: %s", accountID, externalUserID))
	}
	_, err = controller.messengerClient.SendSimpleMessage(externalUserID, fmt.Sprintf("Hello, %s %s, %s", user.FirstName, user.LastName, msg.Text))

	if err != nil {
		return err
	}

	return nil
}

func (controller *SuchefController) fetchUser(accountID int64, externalUserID string) (user *models.User, err error) {
	user, err = controller.dataProvider.FetchUser(accountID, externalUserID)
	if err != nil {
		return nil, err
	}
	if (user != nil) {
		return user, nil
	}

	profile, err := controller.messengerClient.GetProfile(externalUserID)
	if err != nil {
		return nil, err
	}
	_, err = controller.dataProvider.CreateUser(accountID, externalUserID, profile.FirstName, profile.LastName, profile.Gender, profile.ProfilePicture, profile.Locale, profile.Timezone)

	if err != nil {
		return nil, err
	}
	return controller.dataProvider.FetchUser(accountID, externalUserID)
}
