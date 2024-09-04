package main

import (
	"block-game-go/board"
	// "block-game-go/piece"
	"fmt"
)

func main() {
	board.Evaluate()

	fmt.Println(board.ToString())
}
