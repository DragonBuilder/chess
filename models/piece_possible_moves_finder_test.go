package models

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPawnNextMovesFinder_Find_PawnFacingNorth(t *testing.T) {
	b := NewBoard()

	pawn := NewPawn(true)
	pawn.PutOn(b.CellAt("D3"))

	finder := PawnNextMovesFinder{}
	got := finder.Find(pawn)

	assert.Equal(t, 1, len(got), "pawns can only have one possible move at a time")
	assert.True(t, reflect.DeepEqual(b.CellAt("D4"), got[0]), "north facing pawn should return the immediate north neighbour cell")
}

func TestPawnNextMovesFinder_Find_PawnFacingSouth(t *testing.T) {
	b := NewBoard()

	pawn := NewPawn(false)
	pawn.PutOn(b.CellAt("H8"))

	finder := PawnNextMovesFinder{}
	got := finder.Find(pawn)

	assert.Equal(t, 1, len(got), "pawns can only have one possible move at a time")
	assert.True(t, reflect.DeepEqual(b.CellAt("H7"), got[0]), "south facing pawn should return the immediate south neighbour cell")
}

func TestPawnNextMovesFinder_Find_PawnOnTheEdgeCellItFaces(t *testing.T) {
	b := NewBoard()

	southFacingPawn := NewPawn(false)
	southFacingPawn.PutOn(b.CellAt("H1"))

	finder := PawnNextMovesFinder{}
	got1 := finder.Find(southFacingPawn)

	assert.Equal(t, 0, len(got1), "pawn on the last cell in the direction it's facing doesn't have any moves")

	northFacingPawn := NewPawn(true)
	northFacingPawn.PutOn(b.CellAt("A8"))

	got2 := finder.Find(northFacingPawn)

	assert.Equal(t, 0, len(got2), "pawn on the last cell in the direction it's facing doesn't have any moves")
}
