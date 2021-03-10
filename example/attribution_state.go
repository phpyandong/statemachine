package main

type AttributionState struct {
	Code string
	Position int16
	Name	 string
}

func (m *AttributionState) GetCode() (string, error) {
	return m.Code,nil
}

func (m *AttributionState) GetPosition() int16 {
	return m.Position
}







