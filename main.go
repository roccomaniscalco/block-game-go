package main

import (
	"block-game-go/board"
	"block-game-go/piece"
	"fmt"
)

func main() {
	piece := piece.Piece{Shape: "1x1"}

	board.PlacePattern(piece.GetPattern(), board.Coords{0, 0})
	board.PlacePattern(piece.GetPattern(), board.Coords{0, 1})
	board.PlacePattern(piece.GetPattern(), board.Coords{0, 2})
	board.PlacePattern(piece.GetPattern(), board.Coords{1, 0})
	board.PlacePattern(piece.GetPattern(), board.Coords{1, 1})
	board.PlacePattern(piece.GetPattern(), board.Coords{1, 2})
	board.PlacePattern(piece.GetPattern(), board.Coords{2, 0})
	board.PlacePattern(piece.GetPattern(), board.Coords{2, 1})
	board.PlacePattern(piece.GetPattern(), board.Coords{2, 2})

	completedSquares := board.EvaluateSquares()
	fmt.Println(completedSquares)

	err := board.PlacePattern(piece.GetPattern(), board.Coords{5, 5})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(board.ToString())
}
