package configration

import "github.com/statemachine/engine/core"

type IStateMachineConfiguration interface {
	state(core.IState) (StateConfiguration error)
	finalState(state core.IState) StateConfiguration
	getInitialState()	core.IState
	getStateMachineConfiguration() []StateConfiguration
}