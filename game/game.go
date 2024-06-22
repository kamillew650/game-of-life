package game

import "errors"

type GameError struct {
	Message string
}

type Game struct {
	Cells         [][]bool
	height, width int
}

func CreateGame(height, width int) *Game {
	game := new(Game)
	cells := make([][]bool, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]bool, width)
	}
	game.Cells = cells
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

	return game.Cells[heightIndex][widthIndex], nil
}

func (game *Game) SetCellValue(heightIndex int, widthIndex int, value bool) error {
	if (heightIndex+1) > game.height || heightIndex < 0 {
		return errors.New("height index is out of scope")
	}

	if (widthIndex+1) > game.width || widthIndex < 0 {
		return errors.New("width index is out of scope")
	}

	game.Cells[heightIndex][widthIndex] = value

	return nil
}
