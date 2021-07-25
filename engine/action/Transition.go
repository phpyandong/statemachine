package action

import (
	"github.com/statemachine/engine/context"
	"github.com/statemachine/engine/core"
)
type ITranstion interface {
	perform( core.IEvent, ITranstion, context.IGlobalContext)
}
type Transition struct{
	SourceState core.IState
	Destination core.IState
	Guard       IGuard
	Action      IAction
}
//new Transition
func TransitionInstance(action IAction,
	guard IGuard,destination core.IState,sourceState core.IState) Transition {
	transition := Transition{
		Action : action,
		Guard :guard,
		Destination :destination,
		SourceState:sourceState,
	}
	return transition
}
func TransitionInstance2(sourceState core.IState,
	destination core.IState,
	triggers []core.IEvent ) EventTransition {

	transition := EventTransition{
			Transition: Transition{
				SourceState: sourceState,
				Destination: destination,
			},
			Events:triggers,
	}
	return transition
}

type EventTransition struct {
	Transition
	Events []core.IEvent
}
func (eventTransition *EventTransition) perform(event core.Event,transition ITranstion,globalContext context.IGlobalContext)  {
	eventTransition.Action.Perform(event,transition,globalContext)
}
