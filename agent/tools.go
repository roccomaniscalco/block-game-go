package agent

import (
	"block-game-go/game"
	"fmt"
)

type model struct {
	board  game.Board
	pieces game.Pieces
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
	m.pieces.Use(pieceI)

	return nil
}
