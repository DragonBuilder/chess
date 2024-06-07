package models

import (
	"strconv"
)

type board struct {
	cells map[string]*cell
}

func NewBoard() *board {
	b := &board{
		cells: make(map[string]*cell),
	}

	b.populateCells()
	b.linkCells()

	return b
}

func (b *board) CellAt(pos string) *cell {
	return b.cells[pos]
}

func (b *board) Put(piece *Piece, pos string) {
	cell := b.CellAt(pos)
	piece.PutOn(cell)
}

func (b *board) populateCells() {
	for i := 1; i <= 8; i++ {
		for j := 'A'; j <= 'H'; j++ {
			position := string(j) + strconv.Itoa(i)
			b.cells[position] = &cell{
				pos:        position,
				neighbours: make(map[Direction]*cell),
			}
		}
	}
}

func (b *board) linkCells() {
	for _, cell := range b.cells {
		b.linkNorth(cell)
		b.linkNorthEast(cell)
		b.linkEast(cell)
		b.linkSouthEast(cell)
		b.linkSouth(cell)
		b.linkSouthWest(cell)
		b.linkWest(cell)
		b.linkNorthWest(cell)
	}
}

func (b *board) linkNorth(c *cell) {
	col := c.pos[0]
	row := c.pos[1]
	northPos := string(col) + string(row+1)

	if northCell, ok := b.cells[northPos]; ok {
		c.neighbours[NORTH] = northCell
	}
}

func (b *board) linkNorthEast(c *cell) {
	col := c.pos[0]
	row := c.pos[1]

	nwPos := string(col+1) + string(row+1)
	if nwCell, ok := b.cells[nwPos]; ok {
		c.neighbours[NORTH_EAST] = nwCell
	}
}

func (b *board) linkEast(c *cell) {
	col := c.pos[0]
	row := c.pos[1]

	neighbourPos := string(col+1) + string(row)
	if neighbourCell, ok := b.cells[neighbourPos]; ok {
		c.neighbours[EAST] = neighbourCell
	}
}

func (b *board) linkSouthEast(c *cell) {
	col := c.pos[0]
	row := c.pos[1]

	neighbourPos := string(col+1) + string(row-1)
	if neighbourCell, ok := b.cells[neighbourPos]; ok {
		c.neighbours[SOUTH_EAST] = neighbourCell
	}
}

func (b *board) linkSouth(c *cell) {
	col := c.pos[0]
	row := c.pos[1]

	neighbourPos := string(col) + string(row-1)
	if neighbourCell, ok := b.cells[neighbourPos]; ok {
		c.neighbours[SOUTH] = neighbourCell
	}
}

func (b *board) linkSouthWest(c *cell) {
	col := c.pos[0]
	row := c.pos[1]

	neighbourPos := string(col-1) + string(row-1)
	if neighbourCell, ok := b.cells[neighbourPos]; ok {
		c.neighbours[SOUTH_WEST] = neighbourCell
	}
}

func (b *board) linkWest(c *cell) {
	col := c.pos[0]
	row := c.pos[1]

	neighbourPos := string(col-1) + string(row)
	if neighbourCell, ok := b.cells[neighbourPos]; ok {
		c.neighbours[WEST] = neighbourCell
	}
}

func (b *board) linkNorthWest(c *cell) {
	col := c.pos[0]
	row := c.pos[1]

	neighbourPos := string(col-1) + string(row+1)
	if neighbourCell, ok := b.cells[neighbourPos]; ok {
		c.neighbours[NORTH_WEST] = neighbourCell
	}
}
