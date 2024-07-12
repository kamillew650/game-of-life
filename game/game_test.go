package game

import (
	"testing"
)

// TestCreateGame calls greetings.CreateGame and check if cells are created with correct height and width
func TestCreateGame(t *testing.T) {
	const CELL_HEIGHT = 5
	const CELL_WIDTH = 8

	game, err := CreateGame(CELL_HEIGHT, CELL_WIDTH)

	if err != nil {
		t.Errorf(`Game parameters are to low`)
	}

	currentArea := game.getCurrentArea()

	if currentArea == nil {
		t.Errorf(`Game fail to create current area`)
	}
}
