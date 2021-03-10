package core
type IState interface {
	GetCode() (string,error)
	GetPosition() int16
}

type State struct {
	Code string
	Position int16
}
