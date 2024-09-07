package main

import (
	"block-game-go/board"
	// "block-game-go/piece"
	"fmt"
)

func main() {
	b := board.Board{
		Grid: [9][9]bool{
			{false, false, false, true, true,  true,  true,  false, false},
			{false, false, false, true, true,  true,  true,  false, false},
			{false, false, false, true, true,  true,  true,  false, false},
			{false, false, false, true, false, false, false, false, false},
			{false, false, false, true, false, false, false, false, false},
			{false, false, false, true, false, false, false, false, false},
			{true, 	true,  true, 	true, true,  true,  true,  true,  true},
			{false, false, false, true, false, false, false, false, false},
			{false, false, false, true, false, false, false, true,  false},
		},
	}

	fmt.Println(b.ToString())
	b.Evaluate()
	fmt.Println(b.ToString())
}
