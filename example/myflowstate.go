package main

type MyFlowState struct {
	Code string
	Position int16}

func (m *MyFlowState) GetCode() (string, error) {
	return m.Code,nil
}

func (m *MyFlowState) GetPosition() int16 {
	return m.Position
}

