package models

const dimension = 8

type board struct {
	cells [][]int
}

func NewBoard() *board {
	cells := make([][]int, dimension)

	for i := range cells {
		cells[i] = make([]int, dimension)
	}

	return &board{
		cells: cells,
	}
}
