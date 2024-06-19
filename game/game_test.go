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
