package models

type PieceNextMovePossibilitiesFinder interface {
	Find(piece *Piece) []*Cell
}

type PawnNextMovesFinder struct{}

func (p PawnNextMovesFinder) Find(piece *Piece) []*Cell {
	result := make([]*Cell, 0)
	if nc := piece.CurrentCell().NeighbourAt(piece.Facing()); nc != nil {
		result = append(result, nc)
	}
	return result
}

type KingNextMovesFinder struct{}

func (p KingNextMovesFinder) Find(piece *Piece) []*Cell {
	result := make([]*Cell, 0)
	if nc := piece.CurrentCell().NeighbourAt(NORTH); nc != nil {
		result = append(result, nc)
	}
	if nc := piece.CurrentCell().NeighbourAt(NORTH_EAST); nc != nil {
		result = append(result, nc)
	}
	if nc := piece.CurrentCell().NeighbourAt(EAST); nc != nil {
		result = append(result, nc)
	}
	if nc := piece.CurrentCell().NeighbourAt(SOUTH_EAST); nc != nil {
		result = append(result, nc)
	}
	if nc := piece.CurrentCell().NeighbourAt(SOUTH); nc != nil {
		result = append(result, nc)
	}
	if nc := piece.CurrentCell().NeighbourAt(SOUTH_WEST); nc != nil {
		result = append(result, nc)
	}
	if nc := piece.CurrentCell().NeighbourAt(WEST); nc != nil {
		result = append(result, nc)
	}
	if nc := piece.CurrentCell().NeighbourAt(NORTH_WEST); nc != nil {
		result = append(result, nc)
	}
	return result
}

type QueenNextMovesFinder struct{}

func (p QueenNextMovesFinder) Find(piece *Piece) []*Cell {
	result := make([]*Cell, 0)

	result = append(result, cellsTillEdgeInDirection(NORTH, piece.CurrentCell())...)
	result = append(result, cellsTillEdgeInDirection(NORTH_EAST, piece.CurrentCell())...)
	result = append(result, cellsTillEdgeInDirection(EAST, piece.CurrentCell())...)
	result = append(result, cellsTillEdgeInDirection(SOUTH_EAST, piece.CurrentCell())...)
	result = append(result, cellsTillEdgeInDirection(SOUTH, piece.CurrentCell())...)
	result = append(result, cellsTillEdgeInDirection(SOUTH_WEST, piece.CurrentCell())...)
	result = append(result, cellsTillEdgeInDirection(WEST, piece.CurrentCell())...)
	result = append(result, cellsTillEdgeInDirection(NORTH_WEST, piece.CurrentCell())...)

	return result
}

func cellsTillEdgeInDirection(dir Direction, fromCell *Cell) []*Cell {
	result := make([]*Cell, 0)
	current := fromCell
	for current.NeighbourAt(dir) != nil {
		current = current.NeighbourAt(dir)
		result = append(result, current)
	}
	return result
}
