package game

import "errors"

type GameError struct {
	Message string
}

type Game struct {
	cells         [][]bool
	height, width int
}

func CreateGame(height, width int) *Game {
	game := new(Game)
	cells := make([][]bool, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]bool, width)
	}
	game.cells = cells
	game.height = height
	game.width = width
	return game
}

func (game *Game) GetCellValue(heightIndex int, widthIndex int) (bool, error) {
	if (heightIndex+1) > game.height || heightIndex < 0 {
		return false, errors.New("height index is out of scope")
	}

	if (widthIndex+1) > game.width || widthIndex < 0 {
		return false, errors.New("width index is out of scope")
	}

	return game.cells[heightIndex][widthIndex], nil
}

func (game *Game) SetCellValue(heightIndex int, widthIndex int, value bool) error {
	if (heightIndex+1) > game.height || heightIndex < 0 {
		return errors.New("height index is out of scope")
	}

	if (widthIndex+1) > game.width || widthIndex < 0 {
		return errors.New("width index is out of scope")
	}

	game.cells[heightIndex][widthIndex] = value

	return nil
}

func (game *Game) GetCells() [][]bool {
	return game.cells
}
