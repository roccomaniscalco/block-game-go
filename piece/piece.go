package piece

import (
	"math/rand"
)

type Piece struct {
	Grid [][]bool
}

var shapes = [][][]int{
	{
		{1},
	},
	{
		{1, 1},
	},
	{
		{1, 1},
		{1, 1},
	},
	{
		{0, 1},
		{1, 0},
	},
	{
		{1, 0},
		{1, 1},
	},
	{
		{1, 1, 1},
	},
	{
		{1, 0, 0},
		{1, 1, 1},
	},
	{
		{1, 1, 1},
		{0, 1, 0},
		{0, 1, 0},
	},
	{
		{1, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	},
	{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	},
	{
		{1, 1, 1, 1},
	},
	{
		{1, 1, 1, 1, 1},
	},
}

func RandomPiece() Piece {
	randomShape := shapes[rand.Intn(len(shapes))]
	piece := Piece{Grid: convertIntsToBools(randomShape)}

	randomRotationCount := rand.Intn(4)
	for i := 0; i < randomRotationCount; i++ {
		piece.rotate90Deg()
	}

	return piece
}

func convertIntsToBools(shape [][]int) [][]bool {
	grid := make([][]bool, len(shape))
	for rowI := range grid {
		grid[rowI] = make([]bool, len(shape[rowI]))
		for colI := range grid[rowI] {
			grid[rowI][colI] = shape[rowI][colI] == 1
		}
	}
	return grid
}

func (p *Piece) Width() int {
	return len(p.Grid[0])
}

func (p *Piece) Height() int {
	return len(p.Grid)
}

func (p *Piece) rotate90Deg() {
	rowCount := len(p.Grid)
	colCount := len(p.Grid[0])

	rotated := make([][]bool, colCount)
	for rotRowI := range rotated {
		rotated[rotRowI] = make([]bool, rowCount)
	}

	for rowI := range p.Grid {
		for colI := range p.Grid[rowI] {
			rotated[colI][rowCount-1-rowI] = p.Grid[rowI][colI]
		}
	}

	p.Grid = rotated
}

func (p *Piece) ToString() string {
	str := ""
	for rowI := range p.Grid {
		for colI := range p.Grid[rowI] {
			if p.Grid[rowI][colI] {
				str += "▣"
			} else {
				str += " "
			}
			if colI < len(p.Grid[rowI])-1 {
				str += " "
			}
		}
		if rowI < len(p.Grid)-1 {
			str += "\n"
		}
	}
	return str
}
