package models

type Piece struct {
	kind                PieceKind
	facing              Direction
	currentCell         *cell
	nextMoveValidCells  []*cell
	possibilitiesFinder PieceNextMovePossibilitiesFinder
}

func (p *Piece) CurrentCell() *cell {
	return p.currentCell
}

func (p *Piece) Facing() Direction {
	return p.facing
}

func (p *Piece) PutOn(cell *cell) {
	p.currentCell = cell
	cell.Piece = p
	p.nextMoveValidCells = p.possibilitiesFinder.Find(p)
}

func (p *Piece) NextPossibleMoves() []*cell {
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
