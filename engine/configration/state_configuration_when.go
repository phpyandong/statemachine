package configration

import (
	"github.com/statemachine/engine/core"
)

type StateConfigurationWhenSOB struct {
	
}

func (s StateConfigurationWhenSOB) MoveTo(target core.IState) (StateConfiguration, error) {
	panic("implement me")
}

