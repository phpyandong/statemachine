package configration

import (
	"github.com/statemachine/engine/action"
	"github.com/statemachine/engine/core"
	. "time"
)

type IStateConfigurationSOB interface{
	When(events ...core.IEvent) (StateConfigurationWhenSOB)
	MoveAfter(duration Duration,target core.IState) (StateConfiguration,error)
	MoveAfterWithAction(duration Duration, target core.IState, action action.IAction) (StateConfiguration,error)
	MoveAfterIf(duration Duration, target core.IState, action action.IAction,guard action.IGuard) (StateConfiguration,error)
	OnEntry(action action.IAction)
	OnEntryIf(guard action.IGuard,action action.IAction)
	OnExit(action action.IAction)
	OnExitIf(guard action.Guard,action action.IAction)
}