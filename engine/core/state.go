package core
type State struct {
	Code string
	Position int16
}
type IState interface {
	GetCode() (string,error)
	GetPosition() int16
}