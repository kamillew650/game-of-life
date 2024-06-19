package main

import (
	"fmt"
	"game-of-life/game"
)

func main() {
	game := game.CreateGame(4, 2)
	fmt.Println(game.Cells)
}
