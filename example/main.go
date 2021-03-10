package main

import (
	"github.com/statemachine/engine/configration"
	"github.com/statemachine/engine/core"
)

func main()  {

}

func getConfig() configration.IStateMachineConfiguration{
	var  mystate1 AttributionState = AttributionState{Code:"S_1",Position:1}//归属状态实体
	//var  mystate2 AttributionState = AttributionState{Code:"S_2",Position:2}
	//var  mystate3 AttributionState = AttributionState{Code:"S_3",Position:3}

	var myStateA MyFlowState  = MyFlowState{Code:"F_A", Position:0}
	//var myStateB MyFlowState  = MyFlowState{Code:"F_B", Position:1}
	//var myStateC MyFlowState  = MyFlowState{Code:"F_C", Position:2}
	//var myStateD MyFlowState  = MyFlowState{Code:"F_D", Position:3}

	var myEvent core.IEvent = EventNode{Code: "E_1"}
	var stateMechineConfiguration  configration.IStateMachineConfiguration
	var stateConfigure 				configration.StateConfiguration
	var stateConfigurationWhen	configration.StateConfigurationWhenSOB
	//var err	error
	//初始化状态机配置对象
	stateMechineConfiguration = &configration.StateMachineConfiguration{}
	//根据当前状态获取当前的stateConfig
	stateConfigure,_ = stateMechineConfiguration.State(&mystate1)
	//根据stateConfig 配置当前的事件，
	stateConfigurationWhen = stateConfigure.When(myEvent)
	stateConfigurationWhen.MoveTo(&myStateA)
	return stateMechineConfiguration
}