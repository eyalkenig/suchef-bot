package states

import(
	. "github.com/eyalkenig/suchef-bot/server/interaction/inputs"
)

type IState interface{
	Act() (err error)
	Next(input IStateInput) (nextState IState, err error)
	GetNextStage() (IState, error)
	ID() int64
}

type IStateDataProvider interface {

}

type IStateFactory interface {
	GetState(stateID int64) (state IState, err error)
	GetInitialState() (state IState)
}
