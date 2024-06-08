package models

type Direction int

const (
	NORTH Direction = iota
	NORTH_EAST
	EAST
	SOUTH_EAST
	SOUTH
	SOUTH_WEST
	WEST
	NORTH_WEST
)

type PieceKind string

const (
	PAWN  PieceKind = "Pawn"
	KING  PieceKind = "King"
	QUEEN PieceKind = "Queen"
)
