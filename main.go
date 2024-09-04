package main

import (
	"block-game-go/board"
	// "block-game-go/piece"
	"fmt"
)

func main() {
	cells, count := board.Evaluate()
	fmt.Printf("count: %d\n", count)
	fmt.Println(board.ToString(cells))
}
