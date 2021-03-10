package main

import "github.com/statemachine/engine/core"

func main()  {
	var  mystate1 AttributionState = AttributionState{Code:"S_1",Position:1}//归属状态实体
	var  mystate2 AttributionState = AttributionState{Code:"S_2",Position:2}
	var  mystate3 AttributionState = AttributionState{Code:"S_3",Position:3}

	var myStateA MyFlowState  = MyFlowState{Code:"A", Position:0}
	var myStateB MyFlowState  = MyFlowState{Code:"B", Position:1}
	var myStateC MyFlowState  = MyFlowState{Code:"C", Position:2}
	var myStateD MyFlowState  = MyFlowState{Code:"D", Position:3}

}
