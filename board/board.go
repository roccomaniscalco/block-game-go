package board

import (
	"errors"
	"fmt"
	"strings"
)

var board [9][9]int

func init() {
	for rowI := range board {
		for colI := range board[rowI] {
			board[rowI][colI] = 0
		}
	}
}

func PlacePattern(pattern [][]int, coords []int) error {
	startX, startY := coords[0], coords[1]

	if (startX < 0 || startY < 0 || startX > 8 || startY > 8) {
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
		isColumnComplete := true
		for colI := range board[rowI] {
			if board[rowI][colI] == 0 {
				isColumnComplete = false
			}
		}
		if isColumnComplete {
			completedRows = append(completedRows, rowI)
		}
	}

	return completedRows
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
