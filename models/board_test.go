package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	assert.NotNil(t, b)
	assert.Equal(t, dimension, len(b.cells), "Board should have %d rows", dimension)

	for _, row := range b.cells {
		assert.Equal(t, dimension, len(row), "Board should have %d columns", dimension)
	}
}

func Test_Board_Put_When_CorrectPositionProvided_Then_AddsThePieceToCorrectCell(t *testing.T) {
	b := NewBoard()

	pawn := Pawn{}

	err := b.Put(pawn, "A1")

	assert.NoError(t, err)
	assert.Equal(t, PAWN, b.cells[0][0])
}

func Test_Board_Put_When_PositionFormatHasLengthOtherThanTwo_Then_ReturnsError(t *testing.T) {
	b := NewBoard()

	pawn := Pawn{}

	err := b.Put(pawn, "A11")
	assert.Error(t, err)
}

func Test_Board_Put_When_PositionFormatFirstCharIsWrong_Then_ReturnsError(t *testing.T) {
	b := NewBoard()

	pawn := Pawn{}

	err := b.Put(pawn, "K5")
	assert.Error(t, err)
}

func Test_Board_Put_When_PositionFormatSecondCharIsWrong_Then_ReturnsError(t *testing.T) {
	b := NewBoard()

	pawn := Pawn{}

	err := b.Put(pawn, "C9")
	assert.Error(t, err)
}
