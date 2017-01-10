package providers

import "github.com/eyalkenig/suchef-bot/server/models"

type IBotDataProvider interface {
	FetchUser(accountID int64, externalUserID string) (user *models.User, err error)
	CreateUser(accountID int64, externalUserID, firstName, lastName, gender, profilePicURL, locale string, timezone int) (userID int64, err error)
	InitState(userID int64, stateID int64) (err error)
	FetchCurrentState(userID int64) (stateID int64, err error)
	SetCurrentState(userID int64, stateID int64) (err error)

	SetUserDiet(userID, dietTypeID int64) (err error)
	SetSensitivity(userID, sensitivityTypeID int64) (err error)
}

type DBConnectionParams struct {
	User     string
	Password string
	Address  string
	DBName   string
}

type IMessengerProvider interface {
	SendImage(externalUserID, imageURL string) (err error)
	SendSimpleMessage(externalUserID, text string) (err error)
	SendQuickReplyMessage(externalUserID, text string, quickReplies map[string]string) (err error)
	SendGenericTemplate(externalUserID string, titleToPhotoURL map[string]string) (err error)
}
