package piece

var shapes = map[string][][]int{
	"1x1": {{1}},
	"2x2": {{1, 1}, {1, 1}},
}

type Piece struct {
	Shape string
}

func (p *Piece) GetPattern() [][]int {
	return shapes[p.Shape]
}
