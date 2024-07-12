package area

import (
	"testing"
)

// TestCreateGame calls greetings.CreateGame and check if cells are created with correct height and width
func TestCreateArea(t *testing.T) {
	const CELL_HEIGHT = 4
	const CELL_WIDTH = 8

	area := CreateArea(CELL_HEIGHT, CELL_WIDTH, false)

	cellHeight := len(area.GetFields())

	if cellHeight != CELL_HEIGHT {
		t.Errorf(`Cells height value (%d) does not match the initial value (%d)`, cellHeight, CELL_HEIGHT)
	}

	cellWidth := len(area.GetFields()[0])

	if cellWidth != CELL_WIDTH {
		t.Errorf(`Cells width value (%d) does not match the initial value (%d)`, cellWidth, CELL_WIDTH)
	}
}

func TestGetFieldValue(t *testing.T) {
	const CELL_HEIGHT = 4
	const CELL_WIDTH = 8
	const CELL_TEST_HEIGHT_INDEX = 1
	const CELL_TEST_WIDTH_INDEX = 1

	area := CreateArea(CELL_HEIGHT, CELL_WIDTH, false)

	_, error := area.GetFieldValue(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)

	if error != nil {
		t.Errorf(`Getting value within index scope return error`)
	}

	_, heightIndexOutOfScopeError := area.GetFieldValue(-1, CELL_TEST_WIDTH_INDEX)

	if heightIndexOutOfScopeError == nil {
		t.Errorf(`Method allow to get value with heigh index outside the scope`)
	}

	_, widthIndexOutOfScopeError := area.GetFieldValue(CELL_TEST_HEIGHT_INDEX, CELL_WIDTH+3)

	if widthIndexOutOfScopeError == nil {
		t.Errorf(`Method allow to get value with width index outside the scope`)
	}
}

func TestSetFieldValue(t *testing.T) {
	const CELL_HEIGHT = 4
	const CELL_WIDTH = 8
	const CELL_TEST_HEIGHT_INDEX = 1
	const CELL_TEST_WIDTH_INDEX = 1

	area := CreateArea(CELL_HEIGHT, CELL_WIDTH, false)

	heightIndexOutOfScopeError := area.SetFieldValue(-1, CELL_TEST_WIDTH_INDEX, true)

	if heightIndexOutOfScopeError == nil {
		t.Errorf(`Method allow to get value with heigh index outside the scope`)
	}

	widthIndexOutOfScopeError := area.SetFieldValue(CELL_TEST_HEIGHT_INDEX, CELL_WIDTH+3, true)

	if widthIndexOutOfScopeError == nil {
		t.Errorf(`Method allow to get value with width index outside the scope`)
	}

	setValueError := area.SetFieldValue(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX, true)

	if setValueError != nil {
		t.Errorf(`Setting cell value within index scope return error`)
	}

	cellValue, getValueError := area.GetFieldValue(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)

	if cellValue != true {
		t.Errorf(`Setting cell value failed`)
	}

	if getValueError != nil {
		t.Errorf(`Getting cell value within index scope return error`)
	}
}

func TestCalculateNextStateForInactiveField(t *testing.T) {
	const CELL_HEIGHT = 3
	const CELL_WIDTH = 3
	const CELL_TEST_HEIGHT_INDEX = 1
	const CELL_TEST_WIDTH_INDEX = 1

	area := CreateArea(CELL_HEIGHT, CELL_WIDTH, false)

	nextState, err := area.CalculateNextState(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)
	if err != nil {
		t.Errorf(`Calculating next state should not return error`)
	}

	if nextState {
		t.Errorf(`Next state should be equal false`)
	}

	err = area.SetFieldValue(0, 0, true)
	if err != nil {
		t.Errorf(`Height or width index outside the scope`)
	}
	err = area.SetFieldValue(2, 1, true)
	if err != nil {
		t.Errorf(`Height or width index outside the scope`)
	}
	err = area.SetFieldValue(2, 2, true)
	if err != nil {
		t.Errorf(`Height or width index outside the scope`)
	}

	nextState, err = area.CalculateNextState(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)
	if err != nil {
		t.Errorf(`Calculating next state should not return error`)
	}

	if !nextState {
		t.Errorf(`Next state should be equal true`)
	}

	err = area.SetFieldValue(0, 1, true)
	if err != nil {
		t.Errorf(`Height or width index outside the scope`)
	}

	nextState, err = area.CalculateNextState(1, 1)
	if err != nil {
		t.Errorf(`Calculating next state should not return error`)
	}

	if nextState {
		t.Errorf(`Next state should be equal false`)
	}
}

func TestCalculateNextStateForActiveField(t *testing.T) {
	const CELL_HEIGHT = 3
	const CELL_WIDTH = 3
	const CELL_TEST_HEIGHT_INDEX = 1
	const CELL_TEST_WIDTH_INDEX = 1

	area := CreateArea(CELL_HEIGHT, CELL_WIDTH, false)

	err := area.SetFieldValue(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX, true)

	nextState, err := area.CalculateNextState(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)
	if err != nil {
		t.Errorf(`Calculating next state should not return error`)
	}

	if nextState {
		t.Errorf(`Next state should be equal false`)
	}

	err = area.SetFieldValue(0, 0, true)
	if err != nil {
		t.Errorf(`Height or width index outside the scope`)
	}
	err = area.SetFieldValue(2, 1, true)
	if err != nil {
		t.Errorf(`Height or width index outside the scope`)
	}

	nextState, err = area.CalculateNextState(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)
	if err != nil {
		t.Errorf(`Calculating next state should not return error`)
	}

	if !nextState {
		t.Errorf(`Next state should be equal true`)
	}

	err = area.SetFieldValue(2, 2, true)
	if err != nil {
		t.Errorf(`Height or width index outside the scope`)
	}

	nextState, err = area.CalculateNextState(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)
	if err != nil {
		t.Errorf(`Calculating next state should not return error`)
	}

	if !nextState {
		t.Errorf(`Next state should be equal true`)
	}

	err = area.SetFieldValue(0, 1, true)
	if err != nil {
		t.Errorf(`Height or width index outside the scope`)
	}

	nextState, err = area.CalculateNextState(1, 1)
	if err != nil {
		t.Errorf(`Calculating next state should not return error`)
	}

	if nextState {
		t.Errorf(`Next state should be equal false`)
	}
}
