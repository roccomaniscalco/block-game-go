package gameBoard

import (
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
