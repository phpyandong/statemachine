package configration

import (
	"github.com/statemachine/engine/action"
	"github.com/statemachine/engine/context"
	"github.com/statemachine/engine/core"
	. "time"
)

type StateConfiguration struct {
	State core.IState
	EventTransitionsWithoutGuards map[core.IEvent]action.EventTransition
	EventTransitionsWithGuards map[core.IEvent][]action.EventTransition
}
func (s *StateConfiguration) When(events ...core.IEvent) StateConfigurationWhenSOB {
	return StateConfigurationWhenSOB{
			stateConfiguration: *s,
			events: events,
	}
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
//public State fire(final Event event, final StateConfiguration[] destinationsConfiguration,final GlobalContext globalContext)  {
//事件 目的配置
func (s *StateConfiguration) Fire(event core.IEvent,destinationsConfiguration []StateConfiguration, globalContext context.IGlobalContext) {
	var triggerTransitionTarget core.IState = processTriggerTransition(event, destinationsConfiguration,globalContext)
}
func  (s *StateConfiguration) processTriggerTransition(event core.IEvent,destinationsConfiguration []StateConfiguration, globalContext context.IGlobalContext){
	res ,_:= s.EventTransitionsWithoutGuards[event]
	var transitionWithoutGuard action.EventTransition = res
	/**
	  // 是否存在无条件的transition
	    final EventTransition transitionWithoutGuard = eventTransitionsWithoutGuards.get(event);
	    if (transitionWithoutGuard != null) {
	      return performTransition(event, transitionWithoutGuard, destinationConfiguration(transitionWithoutGuard, destinationsConfiguration),globalContext);
	    }
	    // 是否存在有条件的transition
	    else {
	      final List<EventTransition> transitionsWithGuards = eventTransitionsWithGuards.get(event);
	      if (transitionsWithGuards != null) {
	        for (final EventTransition transition : transitionsWithGuards) {
	          if (transition.guard.get().check(event, transition,globalContext)) {
	            return performTransition(event, transition, destinationConfiguration(transition, destinationsConfiguration),globalContext);
	          }
	        }
	        return null;
	      }
	      else {
	        return null; //无效事件
	      }
	    }
	 */
	private State performTransition(final Event event,
		final  Transition transition,
		final StateConfiguration destinationConfiguration,
		final GlobalContext globalContext) {
		transition.perform(event, transition,globalContext);

		transition.destination.ifPresent(destinationState -> {
			//todo 不配置finalState时，destinationConfiguration为空，有啥影响？
			performExitActions(event, transition,globalContext);
			destinationConfiguration.performEntryActions(event, transition,globalContext);
		});

		return transition.destination.isPresent()? transition.destination.get() : transition.source;
	}
}
