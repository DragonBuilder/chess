package errors

import "fmt"

type ErrorUnknownPiece struct {
	PieceName string
}

func (e *ErrorUnknownPiece) Error() string {
	return fmt.Sprintf("unknown piece - %s", e.PieceName)
}

type ErrorUnknownPieceKind struct {
	PieceName string
}

func (e *ErrorUnknownPieceKind) Error() string {
	return fmt.Sprintf("unknown piece kind - %s", e.PieceName)
}

type ErrorInvalidBoardPosition struct {
	Position string
}

func (e *ErrorInvalidBoardPosition) Error() string {
	return fmt.Sprintf("invalid position - %s", e.Position)
}
