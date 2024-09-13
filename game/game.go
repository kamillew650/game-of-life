package game

import (
	"errors"
	"fmt"
	"game-of-life/game/area"
	"time"
)

type Game struct {
	currentArea, nextArea    area.Area
	height, width, iteration int
}

var MIN_HEIGHT, MIN_WIDTH = 5, 5

func CreateGame(height, width int, initialState area.InitialAreaState) (*Game, error) {
	if height < MIN_HEIGHT {
		return nil, errors.New("height is to low")
	}

	if width < MIN_WIDTH {
		return nil, errors.New("width is to low")
	}

	game := new(Game)
	cells := make([][]bool, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]bool, width)
	}
	game.currentArea = *area.CreateArea(height, width, initialState)
	game.height = height
	game.width = width
	game.iteration = 0
	return game, nil
}

func (game *Game) getCurrentArea() *area.Area {
	return &game.currentArea
}

func (game *Game) calculateNextArea() error {
	game.nextArea = *area.CreateArea(game.height, game.width, area.Empty)
	for i := 0; i < game.height; i++ {
		for j := 0; j < game.width; j++ {
			nextState, err := game.currentArea.CalculateNextState(i, j)

			if err != nil {
				return errors.New("fail to calculate next area")
			}

			err = game.nextArea.SetFieldValue(i, j, nextState)

			if err != nil {
				return errors.New("fail to calculate next area")
			}
		}
	}

	game.currentArea = game.nextArea
	return nil
}

func mapBoolToSymbol(value bool) string {
	if value {
		return "X"
	}

	return "O"
}

func (game *Game) displayArea() {
	area := game.currentArea

	fmt.Printf("Iteration: %d \n", game.iteration)
	for i := 0; i < game.height; i++ {
		for j := 0; j < game.width; j++ {
			value, err := area.GetFieldValue(i, j)
			if err != nil {
				return
			}

			fmt.Print(mapBoolToSymbol(value))
			fmt.Print("  ")
		}
		fmt.Print("\n")
	}
}

func (game *Game) StartSimulation(iteration int) error {
	fmt.Print("\033[H\033[2J") // clear console
	for ; game.iteration < iteration; game.iteration++ {
		game.displayArea()
		error := game.calculateNextArea()

		if error != nil {
			return error
		}

		time.Sleep(2 * time.Second)
		fmt.Print("\033[H\033[2J") // clear console
	}

	return nil
}
