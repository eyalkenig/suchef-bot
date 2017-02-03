package server

import (
	"errors"
	"fmt"
	"github.com/eyalkenig/suchef-bot/server/interaction"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	concreteProviders "github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/repositories"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"github.com/eyalkenig/suchef-bot/server/interaction/commands"
)

type SuchefController struct {
	dataProvider     providers.IBotDataProvider
	messengerClient  *messenger.Messenger
	courseRepository repositories.ICourseRepository
}

func NewSuchefController(accountID int64, messengerClient *messenger.Messenger, dataProvider providers.IBotDataProvider) *SuchefController {
	courseRepository := repositories.NewCourseRepository(accountID, dataProvider)
	return &SuchefController{dataProvider: dataProvider, messengerClient: messengerClient, courseRepository: courseRepository}
}

func (controller *SuchefController) Handle(accountID int64, event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) error {
	externalUserID := opts.Sender.ID
	user, err := controller.dataProvider.FetchUser(accountID, externalUserID)

	if err != nil {
		return err
	}

	if user == nil {
		user, err = controller.initUser(accountID, externalUserID)
		if err != nil {
			return err
		}
		if user == nil {
			return errors.New(fmt.Sprintf("user was not found. account id: %d, external user id: %s", accountID, externalUserID))
		}
		return nil
	}

	stateController := controller.getInteractionController(user)
	err = stateController.Handle(msg)

	if err != nil {
		return err
	}

	return nil
}

func (controller *SuchefController) ReloadCourses(accountID int64) {
	controller.courseRepository = repositories.NewCourseRepository(accountID, controller.dataProvider)
}
func (controller *SuchefController) initUser(accountID int64, externalUserID string) (user *models.User, err error) {
	profile, err := controller.messengerClient.GetProfile(externalUserID)
	if err != nil {
		return nil, err
	}
	_, err = controller.dataProvider.CreateUser(accountID, externalUserID, profile.FirstName, profile.LastName, profile.Gender, profile.ProfilePicture, profile.Locale, profile.Timezone)

	if err != nil {
		return nil, err
	}

	user, err = controller.dataProvider.FetchUser(accountID, externalUserID)

	if err != nil {
		return nil, err
	}

	stateController := controller.getInteractionController(user)
	err = stateController.InitUser()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (controller *SuchefController) getInteractionController(user *models.User) interfaces.IInteractionController {
	messengerProvider := concreteProviders.NewFacebookMessengerProvider(controller.messengerClient)
	userContext := context.NewUserContext(user, controller.dataProvider)
	stateMachineController := interaction.NewStateMachineController(messengerProvider, controller.dataProvider, userContext, controller.courseRepository)
	commandsController := commands.NewCommandsController(messengerProvider, controller.dataProvider, userContext)
	return interaction.NewInteractionController(commandsController, stateMachineController)
}
