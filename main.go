package main

import (
	"block-game-go/board"
	// "block-game-go/piece"
	"fmt"
)

func main() {
	b := board.Board{
		Grid: [9][9]int{
			{0, 0, 0, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0, 1, 0},
		},
	}

	fmt.Println(b.ToString())
	b.Evaluate()
	fmt.Println(b.ToString())
}
