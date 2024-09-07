package board

import (
	"errors"
	"fmt"
)

// ANSI escape codes for colors
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

type Board struct {
	Grid [9][9]int
}

type Coords = [2]int

func (b *Board) PlacePattern(pattern [][]int, coords Coords) error {
	startX, startY := coords[0], coords[1]

	if startX < 0 || startY < 0 || startX > 8 || startY > 8 {
		return errors.New("coords must be within range 0-8 inclusive")
	}

	for rowI := range pattern {
		for colI := range pattern[rowI] {
			if startY+rowI > len(b.Grid) || startX+colI > len(b.Grid[rowI]) {
				return errors.New("pattern goes out of bounds")
			}
			if pattern[rowI][colI] == 1 && b.Grid[startY+rowI][startX+colI] == 1 {
				return errors.New("pattern overlaps filled game board tiles")
			}
		}
	}

	for rowI := range pattern {
		for colI := range pattern[rowI] {
			b.Grid[startY+rowI][startX+colI] = pattern[rowI][colI]
		}
	}

	return nil
}

func (b *Board) Evaluate() int {
	completedCells := []Coords{}
	completionCount := 0

	for _, cells := range [][]Coords{b.evaluateRows(), b.evaluateCols(), b.evaluateSquares()} {
		completedCells = append(completedCells, cells...)
		completionCount += len(cells) / 9
	}

	b.removeCells(completedCells)

	return completionCount
}

func (b *Board) evaluateRows() []Coords {
	completedCells := []Coords{}

	for rowI := range b.Grid {
		cells := []Coords{}
		for colI := range b.Grid[rowI] {
			if b.Grid[rowI][colI] == 1 {
				cells = append(cells, Coords{colI, rowI})
			}
		}
		if len(cells) == 9 {
			completedCells = append(completedCells, cells...)
		}
	}

	return completedCells
}

func (b *Board) evaluateCols() []Coords {
	completedCells := []Coords{}

	for colI := range b.Grid[0] {
		cells := []Coords{}
		for rowI := range b.Grid {
			if b.Grid[rowI][colI] == 1 {
				cells = append(cells, Coords{colI, rowI})
			}
		}
		if len(cells) == 9 {
			completedCells = append(completedCells, cells...)
		}
	}

	return completedCells
}

func (b *Board) evaluateSquares() []Coords {
	completedCells := []Coords{}

	// Iterate over each 3x3 square section
	for rowStart := 0; rowStart < 9; rowStart += 3 {
		for colStart := 0; colStart < 9; colStart += 3 {
			cells := []Coords{}

			// Check if all cells in the 3x3 square section are 1s
			for rowI := 0; rowI < 3; rowI++ {
				for colI := 0; colI < 3; colI++ {
					if b.Grid[rowStart+rowI][colStart+colI] == 1 {
						cells = append(cells, Coords{colStart + colI, rowStart + rowI})
					}
				}
			}

			if len(cells) == 9 {
				completedCells = append(completedCells, cells...)
			}
		}
	}

	return completedCells
}

func (b *Board) removeCells(cells []Coords) {
	for rowI := range b.Grid {
		for colI := range b.Grid[rowI] {
			for _, cellToRemove := range cells {
				boardCell := Coords{colI, rowI}
				if cellToRemove == boardCell {
					b.Grid[rowI][colI] = 0
				}
			}
		}
	}
}

func (b *Board) ToString() string {
	str := ""
	for rowI := range b.Grid {
		for colI := range b.Grid[rowI] {

			// isInHighlight := false
			// for _, coord := range highlight {
			// 	if coord[0] == colI && coord[1] == rowI {
			// 		isInHighlight = true
			// 	}
			// }

			// if isInHighlight {
			// 	str += Red + fmt.Sprintf("%d ", b.Grid[rowI][colI]) + Reset
			// } else {
			str += fmt.Sprintf("%d ", b.Grid[rowI][colI])
			// }

		}
		str += "\n"
	}
	return str
}
