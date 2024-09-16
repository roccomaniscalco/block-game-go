package board

import (
	"block-game-go/piece"
	"errors"
)

type Board struct {
	Grid   [9][9]bool
	Score  int
	Streak int
}

type Cell struct {
	RowI int
	ColI int
}

func NewBoard() Board {
	board := Board{
		Score:  0,
		Streak: 0,
	}

	for rowI := range board.Grid {
		for colI := range board.Grid[rowI] {
			board.Grid[rowI][colI] = false
		}
	}

	return board
}

func (b *Board) PlacePiece(piece piece.Piece, start Cell) error {
	rowStart, colStart := start.RowI, start.ColI

	if colStart < 0 || rowStart < 0 || colStart > 8 || rowStart > 8 {
		return errors.New("position must be within range 0-8 inclusive")
	}

	for rowI := range piece.Grid {
		for colI := range piece.Grid[rowI] {
			if rowStart+rowI >= len(b.Grid) || colStart+colI >= len(b.Grid[rowI]) {
				return errors.New("piece goes out of bounds")
			}
			if piece.Grid[rowI][colI] && b.Grid[rowStart+rowI][colStart+colI] {
				return errors.New("piece overlaps filled game board tiles")
			}
		}
	}

	for rowI := range piece.Grid {
		for colI := range piece.Grid[rowI] {
			if piece.Grid[rowI][colI] {
				b.Grid[rowStart+rowI][colStart+colI] = piece.Grid[rowI][colI]
			}
		}
	}

	b.evaluate(piece)

	return nil
}

// Evaluate the board for completed rows, columns, and squares.
// Update the score and streak accordingly.
// Remove the completed cells from the board.
func (b *Board) evaluate(piece piece.Piece) {
	completedCells := []Cell{}
	completionCount := 0

	evaluations := [][]Cell{b.evaluateRows(), b.evaluateCols(), b.evaluateSquares()}
	for _, cells := range evaluations {
		completedCells = append(completedCells, cells...)
		completionCount += len(cells) / 9
	}

	// Update score
	comboBonus := 0
	if completionCount > 1 {
		comboBonus += (completionCount - 1) * 10
	}
	streakBonus := b.Streak * 10
	completionPoints := len(completedCells) * 2

	b.Score += piece.Points() + completionPoints + comboBonus + streakBonus

	// Update streak
	if completionCount > 0 {
		b.Streak += 1
	} else {
		b.Streak = 0
	}

	b.removeCells(completedCells)
}

func (b *Board) evaluateRows() []Cell {
	completedCells := []Cell{}

	for rowI := range b.Grid {
		cells := []Cell{}
		for colI := range b.Grid[rowI] {
			if b.Grid[rowI][colI] {
				cells = append(cells, Cell{RowI: rowI, ColI: colI})
			}
		}
		if len(cells) == 9 {
			completedCells = append(completedCells, cells...)
		}
	}

	return completedCells
}

func (b *Board) evaluateCols() []Cell {
	completedCells := []Cell{}

	for colI := range b.Grid[0] {
		cells := []Cell{}
		for rowI := range b.Grid {
			if b.Grid[rowI][colI] {
				cells = append(cells, Cell{RowI: rowI, ColI: colI})
			}
		}
		if len(cells) == 9 {
			completedCells = append(completedCells, cells...)
		}
	}

	return completedCells
}

func (b *Board) evaluateSquares() []Cell {
	completedCells := []Cell{}

	// Iterate over each 3x3 square section
	for rowStart := 0; rowStart < 9; rowStart += 3 {
		for colStart := 0; colStart < 9; colStart += 3 {
			cells := []Cell{}

			// Check if all cells in the 3x3 square section are 1s
			for rowI := 0; rowI < 3; rowI++ {
				for colI := 0; colI < 3; colI++ {
					if b.Grid[rowStart+rowI][colStart+colI] {
						cell := Cell{RowI: rowStart + rowI, ColI: colStart + colI}
						cells = append(cells, cell)
					}
				}
			}

			if len(cells) == 9 {
				completedCells = append(completedCells, cells...)
			}
		}
	}

	return completedCells
}

func (b *Board) removeCells(cells []Cell) int {
	uniqueCellCount := 0

	for rowI := range b.Grid {
		for colI := range b.Grid[rowI] {
			for _, cellToRemove := range cells {
				boardCell := Cell{RowI: rowI, ColI: colI}
				if cellToRemove == boardCell && b.Grid[rowI][colI] {
					b.Grid[rowI][colI] = false
					uniqueCellCount++
				}
			}
		}
	}

	return uniqueCellCount
}
