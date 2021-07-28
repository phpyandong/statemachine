package engine

import (
	"statemachine/engine/configration"
	"statemachine/engine/context"
	"statemachine/engine/core"
)

type StateMachine struct {
	GlobalContext             context.IGlobalContext
	MachineConfiguration      []configration.StateConfiguration
	CurrentState              core.IState
	CurrentStateConfiguration configration.StateConfiguration
}

func (stateMachine *StateMachine) Fire(trigger core.IEvent) core.IState {
	//获取下一状态
	currentState := stateMachine.CurrentStateConfiguration.Fire(trigger, stateMachine.MachineConfiguration, stateMachine.GlobalContext)
	currentStateConfiguration := stateMachine.MachineConfiguration[currentState.GetPosition()]
	//cleanFormerDelayTransition();
	/**
	 * 开始延时传输
	 */
	//startDelayTransition(currentStateConfiguration);

	return currentState
}
