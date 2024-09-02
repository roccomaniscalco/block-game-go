package main

import (
	"block-game-go/board"
	"block-game-go/piece"
	"fmt"
)

func main() {
	piece := piece.Piece{Shape: "2x2"}
	fmt.Println(piece.GetPattern())
	board.PlacePattern(piece.GetPattern(), []int{1, 1})
	err := board.PlacePattern(piece.GetPattern(), []int{3, 3})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(board.ToString())
}
