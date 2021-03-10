package configration

import "github.com/statemachine/engine/core"

type IStateMachineConfiguration interface {
	State(core.IState) (StateConfiguration, error)
	FinalState(state core.IState)( StateConfiguration,error)
	GetInitialState()	(core.IState,error)
	GetStateMachineConfiguration() map[int16]StateConfiguration
}