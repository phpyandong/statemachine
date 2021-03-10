package configration

import "github.com/statemachine/engine/core"

type IStateConfigurationWhenSOB interface {
	MoveTo(target core.State) (StateConfiguration,error)
}