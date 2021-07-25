package action

import (
	"github.com/statemachine/engine/context"
	"github.com/statemachine/engine/core"
)

type IAction interface {
	Perform( core.Event, ITranstion, context.IGlobalContext)
}
type Action struct {

}