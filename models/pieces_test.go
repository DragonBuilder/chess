package models

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPiece_NextPossibleMoves(t *testing.T) {
	b := NewBoard()

	pawn := NewPawn(true)

	b.Put(pawn, "C5")

	got := pawn.NextPossibleMoves()
	assert.Equal(t, 1, len(got), "pawns can only move one cell in the direction they face.")
	assert.True(t, reflect.DeepEqual(b.CellAt("C6"), got[0]), "north facing pawn can move to C6 from C5")
}
