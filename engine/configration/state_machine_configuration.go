package configration

import (
	"github.com/statemachine/engine/core"
)

type StateMachineConfiguration struct {
	initialState core.IState
	stateConfigurations map[int16]StateConfiguration //positon=>stateconfig
}
/**
	根据状态获取状态的配置，如果没有则创建一个状态配置对象放到 状态机
 */
func (s *StateMachineConfiguration) state(state core.IState) (StateConfiguration ,error) {
	if s.initialState == nil {
		s.initialState = state
	}
	var stateConfiguration StateConfiguration

	if stateConfiguration ,ok := s.stateConfigurations[state.GetPosition()] ; ok==true {
		//return nil,errors.New(fmt.Sprintf("存在重复配置项：%v", state.GetCode()))
	} else {
		stateConfiguration = StateConfiguration{state: state}
		s.stateConfigurations[state.GetPosition()] = stateConfiguration
	}
	return stateConfiguration,nil
}

func (s StateMachineConfiguration) finalState(state core.IState) StateConfiguration {
	panic("implement me")
}

func (s StateMachineConfiguration) getInitialState() core.IState {
	panic("implement me")
}

func (s StateMachineConfiguration) getStateMachineConfiguration() []StateConfiguration {
	panic("implement me")
}
