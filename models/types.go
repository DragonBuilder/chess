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

type PieceKind int

const (
	PAWN PieceKind = 1
)
