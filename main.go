package main

import (
	"block-game-go/gameBoard"
	"block-game-go/gamePiece"
	"fmt"
)

func main() {
	piece := gamePiece.Piece{Shape: "2x2"}
	fmt.Println(piece.GetPattern())
	gameBoard.PlacePattern(piece.GetPattern(), []int{1,1})
	err := gameBoard.PlacePattern(piece.GetPattern(), []int{3,3})
	if (err != nil) {
		fmt.Println(err)
	}

	fmt.Println(gameBoard.ToString())
}
