package models

type cell struct {
	Piece      Piece
	Position   string
	neighbours map[Direction]*cell
}

func (c *cell) NeighbourAt(dir Direction) *cell {
	return c.neighbours[dir]
}
