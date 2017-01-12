package models

import "gopkg.in/maciekmm/messenger-platform-go-sdk.v4"

type User struct {
	ID             int64
	AccountID      int64
	ExternalUserID string
	messenger.Profile
	DietID        *int64
	SensitivityID *int64
}
