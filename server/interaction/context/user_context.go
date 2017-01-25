package context

import (
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/providers"
)

type UserContext struct {
	user       *models.User
	dbProvider providers.IBotDataProvider
}

func NewUserContext(user *models.User, dbProvider providers.IBotDataProvider) *UserContext {
	return &UserContext{user: user, dbProvider: dbProvider}
}

func (context *UserContext) GetID() int64 {
	return context.user.ID
}

func (context *UserContext) GetExternalUserID() string {
	return context.user.ExternalUserID
}

func (context *UserContext) IsMale() bool {
	return context.user.Gender == "male"
}

func (context *UserContext) SetDiet(dietID int64) error {
	return context.dbProvider.SetUserDiet(context.user.ID, dietID)
}

func (context *UserContext) SetSensitivity(sensitivityID int64) error {
	return context.dbProvider.SetSensitivity(context.user.ID, sensitivityID)
}

func (context *UserContext) GetPreferences() (*models.Preference, error) {
	dietID, sensitivityID, err := context.dbProvider.FetchUserPreference(context.user.ID)
	if err != nil {
		return nil, err
	}

	diet, err := models.GetDiet(dietID)

	if err != nil {
		return nil, err
	}

	sensitivity, err := models.GetSensitivity(sensitivityID)

	if err != nil {
		return nil, err
	}

	return &models.Preference{Diet: diet, Sensitivity: sensitivity}, nil
}
