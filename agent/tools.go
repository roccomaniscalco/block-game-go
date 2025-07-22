package agent

import (
	"block-game-go/game"
	"fmt"
	"slices"
)

type model struct {
	board  game.Board
	pieces []game.Piece
}

func (m *model) PlacePiece(pieceI int, rowI int, colI int) error {
	if pieceI < 0 || pieceI >= len(m.pieces) {
		return fmt.Errorf("piece index out of bounds: %d", pieceI)
	}

	piece := m.pieces[pieceI]
	cell := game.Cell{RowI: rowI, ColI: colI}

	if err := m.board.PlacePiece(piece, cell); err != nil {
		return err
	}

	m.board.Evaluate(piece)

	m.pieces = slices.Delete(m.pieces, pieceI, pieceI+1)

	if len(m.pieces) == 0 {
		m.pieces = []game.Piece{game.RandomPiece(), game.RandomPiece(), game.RandomPiece()}
	}

	return nil
}
