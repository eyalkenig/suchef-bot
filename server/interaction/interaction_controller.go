package interaction

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"github.com/eyalkenig/suchef-bot/server/interaction/interfaces"
	"github.com/eyalkenig/suchef-bot/server/interaction/commands"
)

type InteractionController struct {
	stateMachineController interfaces.IStateMachineController
	commandsController interfaces.ICommandsController
}

func NewInteractionController(commandsController interfaces.ICommandsController, stateMachineController interfaces.IStateMachineController) *InteractionController {
	return &InteractionController{stateMachineController: stateMachineController, commandsController: commandsController}
}

func (controller *InteractionController) Handle(message messenger.ReceivedMessage) error {
	if commands.IsCommand(message) {
		return controller.commandsController.Handle(message)
	}
	return controller.stateMachineController.Handle(message)
}

func (controller *InteractionController) InitUser() error {
	return controller.stateMachineController.InitUser()
}
