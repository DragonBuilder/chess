package models

import "github.com/DragonBuilder/chess/errors"

type Piece struct {
	kind                PieceKind
	facing              Direction
	currentCell         *Cell
	nextMoveValidCells  []*Cell
	possibilitiesFinder PieceNextMovePossibilitiesFinder
}

func (p *Piece) CurrentCell() *Cell {
	return p.currentCell
}

func (p *Piece) Facing() Direction {
	return p.facing
}

func (p *Piece) PutOn(cell *Cell) {
	p.currentCell = cell
	cell.Piece = p
	p.nextMoveValidCells = p.possibilitiesFinder.Find(p)
}

func (p *Piece) NextPossibleMoves() []*Cell {
	if p.currentCell == nil {
		panic("cannot determine possible moves on a piece that is not on a cell")
	}
	return p.nextMoveValidCells
}

func NewPawn(northFacing bool) *Piece {
	facing := NORTH
	if !northFacing {
		facing = SOUTH
	}
	return &Piece{
		kind:                PAWN,
		facing:              facing,
		possibilitiesFinder: PawnNextMovesFinder{},
	}
}

func NewKing(northFacing bool) *Piece {
	facing := NORTH
	if !northFacing {
		facing = SOUTH
	}
	return &Piece{
		kind:                KING,
		facing:              facing,
		possibilitiesFinder: KingNextMovesFinder{},
	}
}

func NewQueen(northFacing bool) *Piece {
	facing := NORTH
	if !northFacing {
		facing = SOUTH
	}
	return &Piece{
		kind:                QUEEN,
		facing:              facing,
		possibilitiesFinder: QueenNextMovesFinder{},
	}
}

func NewPiece(kind PieceKind, northFacing bool) (*Piece, error) {
	switch kind {
	case PAWN:
		return NewPawn(northFacing), nil
	case KING:
		return NewKing(northFacing), nil
	case QUEEN:
		return NewQueen(northFacing), nil
	default:
		return nil, &errors.ErrorUnknownPiece{string(kind)}
	}
}
