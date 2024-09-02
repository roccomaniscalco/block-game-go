package gameBoard

import (
	"errors"
	"strings"
)

var gameBoard [9][9]bool

func init() {
	for i := range gameBoard {
		for j := range gameBoard[i] {
			gameBoard[i][j] = false
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
			if startX+i > len(gameBoard) || startY+j > len(gameBoard[i]) {
				return errors.New("pattern goes out of bounds")
			}
			if pattern[i][j] == 1 && gameBoard[startX+i][startY+j] {
				return errors.New("pattern overlaps filled game board tiles")
			}
		}
	}

	for i := range pattern {
		for j := range pattern[i] {
			gameBoard[startX+i][startY+j] = pattern[i][j] == 1
		}
	}

	return nil
}

func ToString() string {
	var builder strings.Builder
	for i := range gameBoard {
		for j := range gameBoard[i] {
			if gameBoard[i][j] {
				builder.WriteString("1 ")
			} else {
				builder.WriteString("0 ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
