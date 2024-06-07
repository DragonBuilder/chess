package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell_NeighbourAt(t *testing.T) {
	b := NewBoard()
	assert.Equal(t, "F6", b.CellAt("F5").NeighbourAt(NORTH).Position())
}

func TestCell_NeighbourAt_when_PositionIsInvalid_then_ReturnNil(t *testing.T) {
	b := NewBoard()
	assert.Nil(t, b.CellAt("H8").NeighbourAt(EAST))
}
