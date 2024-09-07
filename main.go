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

	bCopy := b
	b.Evaluate()
	diff := b.Diff(bCopy)
	fmt.Println(b.ToString(diff))
}
