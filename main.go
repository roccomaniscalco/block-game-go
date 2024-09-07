package main

import (
	"block-game-go/board"
	// "block-game-go/piece"
	"fmt"
)

func main() {
	fmt.Println(board.ToString())
	board.Evaluate()
	fmt.Println(board.ToString())
}
