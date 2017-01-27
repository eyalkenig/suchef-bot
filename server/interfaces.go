package server

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"

	"github.com/eyalkenig/suchef-bot/server/models"
)

type ISuchefController interface {
	Handle(accountID int64, event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) error
}

type ISuchefAdminController interface {
	AddCourse(accountID int64, course *models.Course) error
}
