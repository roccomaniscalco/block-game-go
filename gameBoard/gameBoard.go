package gameBoard

import (
	"errors"
	"fmt"
	"strings"
)

var gameBoard [9][9]int

func init() {
	for i := range gameBoard {
		for j := range gameBoard[i] {
			gameBoard[i][j] = 0
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
			if pattern[i][j] == 1 && gameBoard[startX+i][startY+j] == 1 {
				return errors.New("pattern overlaps filled game board tiles")
			}
		}
	}

	for i := range pattern {
		for j := range pattern[i] {
			gameBoard[startX+i][startY+j] = pattern[i][j]
		}
	}

	return nil
}

func ToString() string {
	var builder strings.Builder
	for i := range gameBoard {
		for j := range gameBoard[i] {
			builder.WriteString(fmt.Sprintf("%d ", gameBoard[i][j]))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
