package models

type Piece interface {
	UID() int
}

type Pawn struct {
}

func (p Pawn) UID() int {
	return PAWN
}
