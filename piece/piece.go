package piece

import (
	"math/rand"
)

var shapes = [][][]int{
	{{1}},
	{{1, 1}, {1, 1}},
}

func GetRandomShape () [][]int {
	randomNumber := rand.Intn(len(shapes))
	return shapes[randomNumber]
}
