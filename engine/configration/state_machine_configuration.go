package configration

import (
	"errors"
	"fmt"
	"github.com/statemachine/engine/core"
)

type StateMachineConfiguration struct {
	initialState core.IState
	stateConfigurations map[int16]StateConfiguration //positon=>stateconfig
}
/**
	根据状态获取状态的配置，如果没有则创建一个状态配置对象放到 状态机
 */
func (s *StateMachineConfiguration) State(state core.IState) (StateConfiguration ,error) {
	if s.initialState == nil {
		s.initialState = state
	}
	if len(s.stateConfigurations) ==0  {
		s.stateConfigurations = make(map[int16]StateConfiguration )
	}
	var stateConfiguration StateConfiguration
	fmt.Printf("position:%+v;configs%+v\n",state.GetPosition(),s.stateConfigurations)
	fmt.Println(state.GetCode())
	position := state.GetPosition()
	if stateConfiguration ,ok := s.stateConfigurations[position] ; ok {
		fmt.Printf("position:%+v;configs%+v state:%+v\n",state.GetPosition(),s.stateConfigurations,state)
		//return stateConfiguration,errors.New(fmt.Sprintf("存在重复配置项：%v", state.GetCode()))
	} else {
		stateConfiguration = StateConfiguration{State: state}
		s.stateConfigurations[position] = stateConfiguration
	}
	fmt.Println("34:",s.stateConfigurations)
	return stateConfiguration,nil
}

func (s *StateMachineConfiguration) FinalState(state core.IState)( StateConfiguration ,error){

	var stateConfiguration StateConfiguration
	fmt.Printf("position:%+v;configs%+v",state.GetPosition(),s.stateConfigurations)

	if stateConfiguration ,ok := s.stateConfigurations[state.GetPosition()] ; ok==true {
		fmt.Printf("position:%+v;configs%+v",state.GetPosition(),s.stateConfigurations)

		//return stateConfiguration,errors.New(fmt.Sprintf("存在重复配置项：%v", state.GetCode()))
	} else {
		stateConfiguration = StateConfiguration{State: state}
		s.stateConfigurations[state.GetPosition()] = stateConfiguration
	}
	return stateConfiguration,nil
}
var MissingStateConfigurationErr = errors.New("initial state is null")
func (s *StateMachineConfiguration) GetInitialState() (core.IState,error) {
	if s.initialState == nil {
		return nil,MissingStateConfigurationErr
	}
	return s.initialState,nil
}
/**
	获取状态机配置中的全部状态配置
 */
func (s *StateMachineConfiguration) GetStateMachineConfiguration() map[int16]StateConfiguration {
	return s.stateConfigurations
}
