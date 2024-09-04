package board

import (
	"errors"
	"fmt"
	"math/rand"
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

var board = [9][9]int{
	{0, 0, 0, 1, 1, 1, 1, 0, 0},
	{0, 0, 0, 1, 1, 1, 1, 0, 0},
	{0, 0, 0, 1, 1, 1, 1, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 1, 0},
}

type Coords = [2]int

func FillEmpty() {
	for rowI := range board {
		for colI := range board[rowI] {
			board[rowI][colI] = 0
		}
	}
}

func FillRandom() {
	for rowI := range board {
		for colI := range board[rowI] {
			board[rowI][colI] = rand.Intn(2)
		}
	}
}

func PlacePattern(pattern [][]int, coords Coords) error {
	startX, startY := coords[0], coords[1]

	if startX < 0 || startY < 0 || startX > 8 || startY > 8 {
		return errors.New("coords must be within range 0-8 inclusive")
	}

	for rowI := range pattern {
		for colI := range pattern[rowI] {
			if startY+rowI > len(board) || startX+colI > len(board[rowI]) {
				return errors.New("pattern goes out of bounds")
			}
			if pattern[rowI][colI] == 1 && board[startY+rowI][startX+colI] == 1 {
				return errors.New("pattern overlaps filled game board tiles")
			}
		}
	}

	for rowI := range pattern {
		for colI := range pattern[rowI] {
			board[startY+rowI][startX+colI] = pattern[rowI][colI]
		}
	}

	return nil
}

func Evaluate() ([]Coords, int) {
	completedCells := []Coords{}
	completionCount := 0

	for _, cells := range [][]Coords{evaluateRows(), evaluateCols(), evaluateSquares()} {
		completedCells = append(completedCells, cells...)
		completionCount += len(cells) / 9
	}

	return completedCells, completionCount
}

func evaluateRows() []Coords {
	completedCells := []Coords{}

	for rowI := range board {
		cells := []Coords{}
		for colI := range board[rowI] {
			if board[rowI][colI] == 1 {
				cells = append(cells, Coords{colI,rowI})
			}
		}
		if len(cells) == 9 {
			completedCells = append(completedCells, cells...)
		}
	}

	return completedCells
}

func evaluateCols() []Coords {
	completedCells := []Coords{}

	for colI := range board[0] {
		cells := []Coords{}
		for rowI := range board {
			if board[rowI][colI] == 1 {
				cells = append(cells, Coords{colI,rowI})
			}
		}
		if len(cells) == 9 {
			completedCells = append(completedCells, cells...)
		}
	}

	return completedCells
}

func evaluateSquares() []Coords {
	completedCells := []Coords{}

	// Iterate over each 3x3 square section
	for rowStart := 0; rowStart < 9; rowStart += 3 {
		for colStart := 0; colStart < 9; colStart += 3 {
			cells := []Coords{}

			// Check if all cells in the 3x3 square section are 1s
			for rowI := 0; rowI < 3; rowI++ {
				for colI := 0; colI < 3; colI++ {
					if board[rowStart+rowI][colStart+colI] == 1 {
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

func ToString(highlight []Coords) string {
	str := ""
	for rowI := range board {
		for colI := range board[rowI] {

			isInHighlight := false
			for _, coord := range highlight {
				if coord[0] == colI && coord[1] == rowI {
					isInHighlight = true
				}
			}

			if isInHighlight {
				str += Red + fmt.Sprintf("%d ", board[rowI][colI]) + Reset
			} else {
				str += fmt.Sprintf("%d ", board[rowI][colI])
			}

		}
		str += "\n"
	}
	return str
}
