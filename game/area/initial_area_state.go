package area

import "errors"

type InitialAreaState int

const (
	Random InitialAreaState = iota
	Light
	Frog
	Empty
)

// String - Creating common behavior - give the type a String function
func (s InitialAreaState) String() string {
	return [...]string{"Random", "Light", "Frog"}[s-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (s InitialAreaState) EnumIndex() int {
	return int(s)
}

func StringToInitialAreaState(s string) (InitialAreaState, error) {
	switch s {
	case "r":
		return Random, nil
	case "l":
		return Light, nil
	case "f":
		return Frog, nil
	default:
		return Random, errors.New("string does not match any inirial area state")
	}
}
