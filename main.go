package main

import (
	"game-of-life/game"
	"log"
)

func main() {
	game, err := game.CreateGame(10, 10)
	if err != nil {
		log.Fatal(1)
	}

	game.StartSimulation(10)
}
