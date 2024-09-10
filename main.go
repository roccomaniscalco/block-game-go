package main

import (
	"block-game-go/board"
	"block-game-go/piece"
	"fmt"
)

func main() {
	b := board.Board{}
	b.Reset()

	piece := piece.RandomPiece()
	b.PlacePattern(piece.Grid, board.Cell{RowI: 0, ColI: 0})

	bCopy := b
	b.Evaluate()
	diff := b.Diff(bCopy)
	fmt.Println(b.ToDiffString(diff))
}
