package models

import (
	"strings"

	"github.com/DragonBuilder/chess/errors"
)

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
	PAWN  PieceKind = "pawn"
	KING  PieceKind = "king"
	QUEEN PieceKind = "queen"
)

func PieceKindFrom(name string) (PieceKind, error) {
	switch strings.ToLower(name) {
	case string(PAWN):
		return PAWN, nil
	case string(KING):
		return KING, nil
	case string(QUEEN):
		return QUEEN, nil
	default:
		return "", &errors.ErrorUnknownPieceKind{name}
	}
}
