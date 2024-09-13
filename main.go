package main

import (
	"flag"
	"game-of-life/game"
	"game-of-life/game/area"
	"log"
	"os"
)

func main() {
	height := flag.Int("height", 10, "height of the area (min and default: 10)")
	width := flag.Int("width", 10, "width of the area (min and default: 10)")
	initialState := flag.String("init", "r", "initial state of simulation, options: random (r), light (l), frog (f)")
	flag.Parse()

	mappedInitialState, error := area.StringToInitialAreaState(*initialState)

	if error != nil {
		log.Fatal(1)
		os.Exit(3)
	}

	game, err := game.CreateGame(*height, *width, mappedInitialState)
	if err != nil {
		log.Fatal(1)
		os.Exit(3)
	}

	if err := game.StartSimulation(10); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
