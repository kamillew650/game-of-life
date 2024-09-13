package area

import (
	"errors"
	"math/rand"
)

type Area struct {
	cells                    [][]bool
	height, width, iteration int
}

func CreateArea(height int, width int, initialState InitialAreaState) *Area {
	area := new(Area)
	cells := make([][]bool, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]bool, width)
	}
	area.cells = cells
	area.height = height
	area.width = width
	area.iteration = 0

	area.setInitialState(initialState)

	return area
}

func (area *Area) setInitialState(initialState InitialAreaState) {
	switch initialState {
	case Random:
		{
			for i := 0; i <= area.height-1; i++ {
				for j := 0; j <= area.width-1; j++ {
					if rand.Intn(10) >= 5 {
						area.SetFieldValue(i, j, true)
					}
				}
			}
		}
	case Frog:
		{
			area.SetFieldValue(1, 2, true)
			area.SetFieldValue(2, 2, true)
			area.SetFieldValue(3, 2, true)
			area.SetFieldValue(2, 3, true)
			area.SetFieldValue(3, 3, true)
			area.SetFieldValue(4, 3, true)
		}

	case Light:
		{
			area.SetFieldValue(1, 2, true)
			area.SetFieldValue(2, 2, true)
			area.SetFieldValue(3, 2, true)
		}
	}
}

func (area *Area) GetFieldValue(heightIndex int, widthIndex int) (bool, error) {
	if (heightIndex+1) > area.height || heightIndex < 0 {
		return false, errors.New("height index is out of scope")
	}

	if (widthIndex+1) > area.width || widthIndex < 0 {
		return false, errors.New("width index is out of scope")
	}

	return area.cells[heightIndex][widthIndex], nil
}

func (area *Area) SetFieldValue(heightIndex int, widthIndex int, value bool) error {
	if (heightIndex+1) > area.height || heightIndex < 0 {
		return errors.New("height index is out of scope")
	}

	if (widthIndex+1) > area.width || widthIndex < 0 {
		return errors.New("width index is out of scope")
	}

	area.cells[heightIndex][widthIndex] = value

	return nil
}

func (area *Area) GetFields() [][]bool {
	return area.cells
}

func (area *Area) CalculateNextState(heightIndex int, widthIndex int) (bool, error) {
	currentCellState, err := area.GetFieldValue(heightIndex, widthIndex)
	if err != nil {
		return false, errors.New("heigh or width index is out of scope")
	}

	var activeCellsAround = 0

	for moveHeightIndex := -1; moveHeightIndex <= 1; moveHeightIndex++ {
		for moveWidthIndex := -1; moveWidthIndex <= 1; moveWidthIndex++ {
			// fmt.Print(heightIndex+moveHeightIndex, widthIndex+moveWidthIndex, "\n")
			if moveHeightIndex != 0 || moveWidthIndex != 0 {
				value, err := area.GetFieldValue(heightIndex+moveHeightIndex, widthIndex+moveWidthIndex)

				if err == nil && value {
					activeCellsAround++
				}
			}
		}
	}

	if !currentCellState && activeCellsAround == 3 {
		return true, nil
	} else if currentCellState && (activeCellsAround == 2 || activeCellsAround == 3) {
		return true, nil
	}

	return false, nil
}
