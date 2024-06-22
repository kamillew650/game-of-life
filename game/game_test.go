package game

import (
	"testing"
)

// TestCreateGame calls greetings.CreateGame and check if cells are created with correct height and width
func TestCreateGame(t *testing.T) {
	const CELL_HEIGHT = 4
	const CELL_WIDTH = 8

	game := CreateGame(CELL_HEIGHT, CELL_WIDTH)

	cellHeight := len(game.Cells)

	if cellHeight != CELL_HEIGHT {
		t.Errorf(`Cells height value (%d) does not match the initial value (%d)`, cellHeight, CELL_HEIGHT)
	}

	cellWidth := len(game.Cells[0])

	if cellWidth != CELL_WIDTH {
		t.Errorf(`Cells width value (%d) does not match the initial value (%d)`, cellWidth, CELL_WIDTH)
	}
}

func TestGetCellValue(t *testing.T) {
	const CELL_HEIGHT = 4
	const CELL_WIDTH = 8
	const CELL_TEST_HEIGHT_INDEX = 1
	const CELL_TEST_WIDTH_INDEX = 1

	game := CreateGame(CELL_HEIGHT, CELL_WIDTH)

	_, error := game.GetCellValue(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)

	if error != nil {
		t.Errorf(`Getting value within index scope return error`)
	}

	_, heightIndexOutOfScopeError := game.GetCellValue(-1, CELL_TEST_WIDTH_INDEX)

	if heightIndexOutOfScopeError == nil {
		t.Errorf(`Method allow to get value with heigh index outside the scope`)
	}

	_, widthIndexOutOfScopeError := game.GetCellValue(CELL_TEST_HEIGHT_INDEX, CELL_WIDTH+3)

	if widthIndexOutOfScopeError == nil {
		t.Errorf(`Method allow to get value with width index outside the scope`)
	}
}

func TestSetCellValue(t *testing.T) {
	const CELL_HEIGHT = 4
	const CELL_WIDTH = 8
	const CELL_TEST_HEIGHT_INDEX = 1
	const CELL_TEST_WIDTH_INDEX = 1

	game := CreateGame(CELL_HEIGHT, CELL_WIDTH)

	heightIndexOutOfScopeError := game.SetCellValue(-1, CELL_TEST_WIDTH_INDEX, true)

	if heightIndexOutOfScopeError == nil {
		t.Errorf(`Method allow to get value with heigh index outside the scope`)
	}

	widthIndexOutOfScopeError := game.SetCellValue(CELL_TEST_HEIGHT_INDEX, CELL_WIDTH+3, true)

	if widthIndexOutOfScopeError == nil {
		t.Errorf(`Method allow to get value with width index outside the scope`)
	}

	setValueError := game.SetCellValue(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX, true)

	if setValueError != nil {
		t.Errorf(`Setting cell value within index scope return error`)
	}

	cellValue, getValueError := game.GetCellValue(CELL_TEST_HEIGHT_INDEX, CELL_TEST_WIDTH_INDEX)

	if cellValue != true {
		t.Errorf(`Setting cell value failed`)
	}

	if getValueError != nil {
		t.Errorf(`Getting cell value within index scope return error`)
	}
}
