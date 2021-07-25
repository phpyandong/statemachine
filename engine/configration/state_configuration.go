package configration

import (
	"github.com/statemachine/engine/action"
	"github.com/statemachine/engine/core"
	. "time"
)

type StateConfiguration struct {
	State core.IState
	EventTransitionsWithoutGuards map[core.IEvent]action.EventTransition
	EventTransitionsWithGuards map[core.IEvent][]action.EventTransition
}
func (s *StateConfiguration) When(events ...core.IEvent) StateConfigurationWhenSOB {
	return StateConfigurationWhenSOB{
			stateConfiguration: *s,
			events: events,
	}
}

func (s *StateConfiguration) MoveAfter(duration Duration, target core.IState) (StateConfiguration, error) {
	panic("implement me")
}

func (s *StateConfiguration) MoveAfterWithAction(duration Duration, target core.IState, action action.IAction) (StateConfiguration, error) {
	panic("implement me")
}

func (s *StateConfiguration) MoveAfterIf(duration Duration, target core.IState, action action.IAction, guard action.IGuard) (StateConfiguration, error) {
	panic("implement me")
}

func (s *StateConfiguration) OnEntry(action action.IAction) {
	panic("implement me")
}

func (s *StateConfiguration) OnEntryIf(guard action.IGuard, action action.IAction) {
	panic("implement me")
}

func (s *StateConfiguration) OnExit(action action.IAction) {
	panic("implement me")
}

func (s *StateConfiguration) OnExitIf(guard action.Guard, action action.IAction) {
	panic("implement me")
}

func NewStateConfiguration() *StateConfiguration {
	return &StateConfiguration{}
}
