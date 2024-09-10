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
	piece := Piece{Grid: convertShapeToGrid(randomShape)}

	randomRotationCount := rand.Intn(4)
	for i := 0; i < randomRotationCount; i++ {
		piece.rotate90Deg()
	}

	return piece
}

func convertShapeToGrid(shape [][]int) [][]bool {
	grid := make([][]bool, len(shape))
	for rowI := range grid {
		grid[rowI] = make([]bool, len(shape[rowI]))
		for colI := range grid[rowI] {
			grid[rowI][colI] = shape[rowI][colI] == 1
		}
	}
	return grid
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
				str += "▣ "
			} else {
				str += "□ "
			}
		}
		str += "\n"
	}
	return str
}
