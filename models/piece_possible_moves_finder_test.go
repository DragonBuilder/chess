package models_test

import (
	"reflect"
	"testing"

	"github.com/DragonBuilder/chess/models"
	"github.com/stretchr/testify/assert"
)

func TestPawnNextMovesFinder_Find_PawnFacingNorth(t *testing.T) {
	b := models.NewBoard()

	pawn := models.NewPawn(true)
	pawn.PutOn(b.CellAt("D3"))

	finder := models.PawnNextMovesFinder{}
	got := finder.Find(pawn)

	assert.Equal(t, 1, len(got), "pawns can only have one possible move at a time")
	assert.True(t, reflect.DeepEqual(b.CellAt("D4"), got[0]), "north facing pawn should return the immediate north neighbour cell")
}

func TestPawnNextMovesFinder_Find_PawnFacingSouth(t *testing.T) {
	b := models.NewBoard()

	pawn := models.NewPawn(false)
	pawn.PutOn(b.CellAt("H8"))

	finder := models.PawnNextMovesFinder{}
	got := finder.Find(pawn)

	assert.Equal(t, 1, len(got), "pawns can only have one possible move at a time")
	assert.True(t, reflect.DeepEqual(b.CellAt("H7"), got[0]), "south facing pawn should return the immediate south neighbour cell")
}

func TestPawnNextMovesFinder_Find_PawnOnTheEdgeCellItFaces(t *testing.T) {
	b := models.NewBoard()

	southFacingPawn := models.NewPawn(false)
	southFacingPawn.PutOn(b.CellAt("H1"))

	finder := models.PawnNextMovesFinder{}
	got1 := finder.Find(southFacingPawn)

	assert.Equal(t, 0, len(got1), "pawn on the last cell in the direction it's facing doesn't have any moves")

	northFacingPawn := models.NewPawn(true)
	northFacingPawn.PutOn(b.CellAt("A8"))

	got2 := finder.Find(northFacingPawn)

	assert.Equal(t, 0, len(got2), "pawn on the last cell in the direction it's facing doesn't have any moves")
}

func TestKingNextMovesFinder_Find(t *testing.T) {
	b := models.NewBoard()

	king := models.NewKing(true)

	b.Put(king, "D5")
	finder := models.KingNextMovesFinder{}
	got := finder.Find(king)

	assert.Equal(t, 8, len(got), "king should have 8 possible moves in all 8 directions")
}

func TestQueenNextMovesFinder_Find(t *testing.T) {
	b := models.NewBoard()

	queen := models.NewQueen(true)

	b.Put(queen, "E4")
	finder := models.QueenNextMovesFinder{}
	got := finder.Find(queen)

	assert.Equal(t, 27, len(got), "queen can move in all 8 directions untill edge")
}
