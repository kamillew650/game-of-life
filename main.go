package main

import (
	"fmt"
	"game-of-life/game"
	"log"
)

func main() {
	game := game.CreateGame(4, 2)
	err := game.SetCellValue(2, 1, true)
	if err != nil {
		log.Fatal(err)
	}

	value, err := game.GetCellValue(2, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
	fmt.Println(game.GetCells())
}
