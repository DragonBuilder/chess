package models

type Cell struct {
	Piece *Piece

	pos        string
	neighbours map[Direction]*Cell
}

func (c *Cell) Position() string {
	return c.pos
}

func (c *Cell) NeighbourAt(dir Direction) *Cell {
	return c.neighbours[dir]
}
