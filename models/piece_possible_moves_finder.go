package models

type PieceNextMovePossibilitiesFinder interface {
	Find(piece *Piece) []*cell
}

type PawnNextMovesFinder struct{}

func (p PawnNextMovesFinder) Find(piece *Piece) []*cell {
	result := make([]*cell, 0)
	if nc := piece.CurrentCell().NeighbourAt(piece.Facing()); nc != nil {
		result = append(result, nc)
	}
	return result
}
