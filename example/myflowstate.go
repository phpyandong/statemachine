package main

type MyFlowState struct {
	Code string
	Position int16}

func (m *MyFlowState) getCode() (string, error) {
	return m.Code,nil
}

func (m *MyFlowState) getPosition() int16 {
	return m.Position
}

