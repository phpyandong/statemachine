package configration

import (
	"fmt"
	"github.com/statemachine/engine/action"
	"github.com/statemachine/engine/core"
)
//private final StateConfiguration configuration;
//private final Event[] triggers;
type StateConfigurationWhenSOB struct {
	stateConfiguration StateConfiguration
	events []core.IEvent
}

func (s *StateConfigurationWhenSOB) MoveTo(target core.IState) (StateConfiguration, error) {
	fmt.Println("\ntarget:",target)
	fmt.Println("\neveentTranstion:",s.stateConfiguration.EventTransitionsWithoutGuards)

	for _,event := range s.events {
		fmt.Println("event",event)
		fmt.Println("new trans:",action.TransitionInstance2(s.stateConfiguration.State,target,s.events ) )
		evenMap := s.stateConfiguration.EventTransitionsWithoutGuards
		if len(evenMap) ==0 {
			evenMap := make(map[core.IEvent]action.EventTransition)
			evenMap[event] = action.TransitionInstance2(s.stateConfiguration.State,target,s.events )
			s.stateConfiguration.EventTransitionsWithoutGuards = evenMap
		}

		//fmt.Printf("eventMap x %+v\n",evenMap)
		//
		//fmt.Printf("eventMap 2 %+v\n",evenMap)

	}
	//fmt.Println("\neveentTranstion new:",s.stateConfiguration.EventTransitionsWithoutGuards)

	return s.stateConfiguration,nil

}

