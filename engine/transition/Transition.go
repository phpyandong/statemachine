package transition

import (
	"github.com/statemachine/engine/action"
	"github.com/statemachine/engine/core"
)
type ITranstion interface {
	perform()
}
type Transition struct{
	destination core.State
	guard		action.Guard
	actionObj 	action.Action
}
