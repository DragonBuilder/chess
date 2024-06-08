package services

import (
	"strings"

	"github.com/DragonBuilder/chess/errors"
	"github.com/DragonBuilder/chess/models"
)

func Put(piece string, position string) ([]string, error) {
	pKind, err := models.PieceKindFrom(piece)
	if err != nil {
		return nil, err
	}

	p, err := models.NewPiece(pKind, true)
	if err != nil {
		return nil, err
	}

	b := models.NewBoard()
	position = strings.ToUpper(position)

	if b.CellAt(position) == nil {
		return nil, &errors.ErrorInvalidBoardPosition{Position: position}
	}

	b.Put(p, position)

	moves := p.NextPossibleMoves()

	movesStr := make([]string, 0)
	for _, m := range moves {
		movesStr = append(movesStr, m.Position())
	}
	return movesStr, nil
}
