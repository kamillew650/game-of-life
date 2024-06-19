package game

type Game struct {
	Cells [][]bool
}

func CreateGame(height, width int) *Game {
	game := new(Game)
	cells := make([][]bool, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]bool, width)
	}
	game.Cells = cells
	return game
}
