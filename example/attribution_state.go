package main

type AttributionState struct {
	Code string
	Position int16
	Name	 string
}

func (m *AttributionState) getCode() (string, error) {
	return m.Code,nil
}

func (m *AttributionState) getPosition() int16 {
	return m.Position
}







