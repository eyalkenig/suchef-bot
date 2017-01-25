package context

import "github.com/eyalkenig/suchef-bot/server/models"

type IUserContext interface {
	GetID() int64
	GetExternalUserID() string
	IsMale() bool

	SetDiet(dietID int64) error
	SetSensitivity(sensitivityID int64) error

	GetPreferences() (*models.Preference, error)
}
