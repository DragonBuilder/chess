package models

type cell struct {
	Piece *Piece

	pos        string
	neighbours map[Direction]*cell
}

func (c *cell) Position() string {
	return c.pos
}

func (c *cell) NeighbourAt(dir Direction) *cell {
	return c.neighbours[dir]
}
