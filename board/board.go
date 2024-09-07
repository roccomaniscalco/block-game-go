package board

import (
	"errors"
)

type Board struct {
	Grid [9][9]bool
}

type Cell struct {
	RowI int
	ColI int
}

func (b *Board) PlacePattern(pattern [][]bool, startCell Cell) error {
	rowStart, colStart := startCell.RowI, startCell.ColI

	if colStart < 0 || rowStart < 0 || colStart > 8 || rowStart > 8 {
		return errors.New("position must be within range 0-8 inclusive")
	}

	for rowI := range pattern {
		for colI := range pattern[rowI] {
			if rowStart+rowI > len(b.Grid) || colStart+colI > len(b.Grid[rowI]) {
				return errors.New("pattern goes out of bounds")
			}
			if pattern[rowI][colI] && b.Grid[rowStart+rowI][colStart+colI] {
				return errors.New("pattern overlaps filled game board tiles")
			}
		}
	}

	for rowI := range pattern {
		for colI := range pattern[rowI] {
			b.Grid[rowStart+rowI][colStart+colI] = pattern[rowI][colI]
		}
	}

	return nil
}

func (b *Board) Evaluate() int {
	completedCells := []Cell{}
	completionCount := 0

	evaluations := [][]Cell{b.evaluateRows(), b.evaluateCols(), b.evaluateSquares()}
	for _, cells := range evaluations {
		completedCells = append(completedCells, cells...)
		completionCount += len(cells) / 9
	}

	b.removeCells(completedCells)

	return completionCount
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

func (b *Board) removeCells(cells []Cell) {
	for rowI := range b.Grid {
		for colI := range b.Grid[rowI] {
			for _, cellToRemove := range cells {
				boardCell := Cell{RowI: rowI, ColI: colI}
				if cellToRemove == boardCell {
					b.Grid[rowI][colI] = false
				}
			}
		}
	}
}

func (b *Board) ToString() string {
	str := ""
	for rowI := range b.Grid {
		for colI := range b.Grid[rowI] {
			if b.Grid[rowI][colI] {
				str += "▣ "
			} else {
				str += "□ "
			}
		}
		str += "\n"
	}
	return str
}
