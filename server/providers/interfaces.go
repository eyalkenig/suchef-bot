package providers

import "github.com/eyalkenig/suchef-bot/server/models"

type IBotDataProvider interface {
	FetchUser(accountID int64, externalUserID string) (user *models.User, err error)
	CreateUser(accountID int64, externalUserID, firstName, lastName, gender, profilePicURL, locale string, timezone int) (userID int64, err error)
}

type DBConnectionParams struct {
	User string
	Password string
	Address string
	DBName string
}