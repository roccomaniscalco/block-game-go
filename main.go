package main

import (
	"block-game-go/board"
	"block-game-go/piece"
	"fmt"
)

func main() {
	piece := piece.Piece{Shape: "1x1"}
	fmt.Println(piece.GetPattern())
	board.PlacePattern(piece.GetPattern(), []int{0, 3})
	board.PlacePattern(piece.GetPattern(), []int{1, 3})
	board.PlacePattern(piece.GetPattern(), []int{2, 3})
	board.PlacePattern(piece.GetPattern(), []int{3, 3})
	board.PlacePattern(piece.GetPattern(), []int{4, 3})
	board.PlacePattern(piece.GetPattern(), []int{5, 3})
	board.PlacePattern(piece.GetPattern(), []int{6, 3})
	board.PlacePattern(piece.GetPattern(), []int{7, 3})
	board.PlacePattern(piece.GetPattern(), []int{8, 3})

	board.PlacePattern(piece.GetPattern(), []int{4, 0})
	board.PlacePattern(piece.GetPattern(), []int{4, 1})
	board.PlacePattern(piece.GetPattern(), []int{4, 2})
	board.PlacePattern(piece.GetPattern(), []int{4, 3})
	board.PlacePattern(piece.GetPattern(), []int{4, 4})
	board.PlacePattern(piece.GetPattern(), []int{4, 5})
	board.PlacePattern(piece.GetPattern(), []int{4, 6})
	board.PlacePattern(piece.GetPattern(), []int{4, 7})
	board.PlacePattern(piece.GetPattern(), []int{4, 8})

	completedRows := board.EvaluateRows()
	completedCols := board.EvaluateCols()
	fmt.Println(completedRows, completedCols)

	err := board.PlacePattern(piece.GetPattern(), []int{2, 3})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(board.ToString())
}
