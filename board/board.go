package board

import (
	"errors"
	"fmt"
	"strings"
)

var board [9][9]int

type Coords = [2]int

func init() {
	for rowI := range board {
		for colI := range board[rowI] {
			board[rowI][colI] = 0
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

func EvaluateRows() []int {
	completedRows := []int{}

	for rowI := range board {
		isRowComplete := true
		for colI := range board[rowI] {
			if board[rowI][colI] == 0 {
				isRowComplete = false
				break
			}
		}
		if isRowComplete {
			completedRows = append(completedRows, rowI)
		}
	}

	return completedRows
}

func EvaluateCols() []int {
	completedCols := []int{}

	for colI := range board[0] {
		isColComplete := true
		for rowI := range board {
			if board[rowI][colI] == 0 {
				isColComplete = false
				break
			}
		}
		if isColComplete {
			completedCols = append(completedCols, colI)
		}
	}

	return completedCols
}

func EvaluateSquares() []Coords {
	completedCells := []Coords{}

	// Iterate over each 3x3 section
	for rowStart := 0; rowStart < 9; rowStart += 3 {
			for colStart := 0; colStart < 9; colStart += 3 {
					cells := []Coords{}

					// Check if all elements in the 3x3 section are 1s
					for rowI := 0; rowI < 3; rowI++ {
							for colI := 0; colI < 3; colI++ {
								if board[rowStart+rowI][colStart+colI] == 1 {
									cells = append(cells, Coords{colStart+colI, rowStart+rowI})
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

func ToString() string {
	var builder strings.Builder
	for rowI := range board {
		for colI := range board[rowI] {
			builder.WriteString(fmt.Sprintf("%d ", board[rowI][colI]))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
