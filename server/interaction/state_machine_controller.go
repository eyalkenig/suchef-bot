package interaction

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"github.com/eyalkenig/suchef-bot/server/interaction/states"
	"github.com/eyalkenig/suchef-bot/server/interaction/inputs"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
)

type StateMachineController struct {
	inputFactory      inputs.IStateInputFactory
	statesFactory     states.IStateFactory
	stateDataProvider providers.IBotDataProvider
	userContext context.IUserContext
}

func NewStateMachineController(messengerProvider providers.IMessengerProvider,
			       stateDataProvider providers.IBotDataProvider,
			       userContext context.IUserContext) *StateMachineController {
	inputFactory := inputs.NewStateInputFactory()
	statesFactory := states.NewStateFactory(messengerProvider, stateDataProvider, userContext)

	return &StateMachineController{inputFactory: inputFactory,
				       stateDataProvider: stateDataProvider,
				       statesFactory: statesFactory,
				       userContext: userContext}
}

func (controller *StateMachineController) InitUser() (err error) {
	initialState := controller.statesFactory.GetInitialState()
	err = initialState.Act()
	if (err != nil) {
		return err
	}
	return controller.stateDataProvider.InitState(controller.userContext.GetID(), initialState.ID())
}

func (controller *StateMachineController) Handle(message messenger.ReceivedMessage) (err error) {
	userID := controller.userContext.GetID()
	stateID, err := controller.stateDataProvider.FetchCurrentState(userID)
	if (err != nil) {
		return err
	}
	state, err := controller.statesFactory.GetState(stateID)
	if (err != nil) {
		return err
	}
	input, err := controller.inputFactory.CastMessageToInput(message)
	if (err != nil) {
		return err
	}

	payload := input.Payload()
	if (payload == inputs.LETS_START_FROM_SCRATCH_INPUT){
		return controller.InitUser()
	}

	nextState, err := state.Next(input)

	if (err != nil) {
		return err
	}
	if (nextState != nil) {
		err = nextState.Act()
		if (err != nil){
			return err
		}

		err = controller.stateDataProvider.SetCurrentState(userID, nextState.ID())
		if (err != nil) {
			return err
		}

		nextStage, err := nextState.GetNextStage()
		if (err != nil) {
			return err
		}

		if nextStage != nil {
			err = nextStage.Act()
			if (err != nil){
				return err
			}

			err = controller.stateDataProvider.SetCurrentState(userID, nextStage.ID())
			if (err != nil) {
				return err
			}
		}
	}
	return nil
}
