# Game of Life in Go

This is a simple console-based implementation of Conway's Game of Life written in Go. The program allows users to specify the grid dimensions and initial state of the game area.

## Requirements

- Go 1.16 or higher

## Usage

To run the Game of Life, use the following command:

```
go run main.go --height [height] --width [width] --init [initial state]
```

### Parameters

- `height`: Integer value for the grid height (default: 10)
- `width`: Integer value for the grid width (default: 10)
- `initial state`: String representing the initial state of the grid (default: random)

The `initial state` parameter should be a string. There are options r (random), l (light), f (frog).

### Examples

Run with default settings (15x15 grid):

```
go run main.go --height 15 --width 15 --init r
```

## Game Rules

The Game of Life follows these rules:

1. Any live cell with fewer than two live neighbors dies (underpopulation).
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies (overpopulation).
4. Any dead cell with exactly three live neighbors becomes a live cell (reproduction).

## License

This project is open source and available under the [MIT License](LICENSE).

## Author

Kamil Lewandowski
