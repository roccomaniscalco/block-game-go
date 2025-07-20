package agent

import (
	"block-game-go/board"
	"block-game-go/piece"
	"fmt"
	"slices"
)

type state struct {
	board  board.Board
	pieces []piece.Piece
}

func (s *state) PlacePiece(pieceI int, rowI int, colI int) error {
	if pieceI < 0 || pieceI >= len(s.pieces) {
		return fmt.Errorf("piece index out of bounds: %d", pieceI)
	}

	p := s.pieces[pieceI]
	c := board.Cell{RowI: rowI, ColI: colI}

	if err := s.board.PlacePiece(p, c); err != nil {
		return err
	}

	s.board.Evaluate(p)

	s.pieces = slices.Delete(s.pieces, pieceI, pieceI+1)

	if len(s.pieces) == 0 {
		s.pieces = []piece.Piece{piece.RandomPiece(), piece.RandomPiece(), piece.RandomPiece()}
	}

	return nil
}
