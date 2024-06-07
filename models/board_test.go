package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	assert.NotNil(t, b)
	assert.Equal(t, 64, len(b.cells), "Board should have %d cells", 64)
}

func TestBoard_CellAt_when_PositionIsValid_then_ReturnsCell(t *testing.T) {
	b := NewBoard()
	assert.NotNil(t, b.CellAt("A1"))
}

func TestBoard_CellAt_when_PositionIsInvalid_then_ReturnsNil(t *testing.T) {
	b := NewBoard()
	assert.Nil(t, b.CellAt("H15"))
}
