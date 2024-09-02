package board

import (
	"errors"
	"fmt"
	"strings"
)

var board [9][9]int

func init() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = 0
		}
	}
}

func PlacePattern(pattern [][]int, coords []int) error {
	startX, startY := coords[0], coords[1]

	if (startX < 1 || startY < 1) {
		return errors.New("coords must be positive")
	}

	for i := range pattern {
		for j := range pattern[i] {
			if startX+i > len(board) || startY+j > len(board[i]) {
				return errors.New("pattern goes out of bounds")
			}
			if pattern[i][j] == 1 && board[startX+i][startY+j] == 1 {
				return errors.New("pattern overlaps filled game board tiles")
			}
		}
	}

	for i := range pattern {
		for j := range pattern[i] {
			board[startX+i][startY+j] = pattern[i][j]
		}
	}

	return nil
}

func ToString() string {
	var builder strings.Builder
	for i := range board {
		for j := range board[i] {
			builder.WriteString(fmt.Sprintf("%d ", board[i][j]))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
