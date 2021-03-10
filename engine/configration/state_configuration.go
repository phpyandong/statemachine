package configration

import (
	. "time"
	"github.com/statemachine/engine/action"
	"github.com/statemachine/engine/core"
)

type StateConfiguration struct{
	state core.IState
}

func (s *StateConfiguration) When(events ...core.IEvent) (StateConfigurationWhenSOB) {
	panic("implement me")
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