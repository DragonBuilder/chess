package models_test

import (
	"testing"

	"github.com/DragonBuilder/chess/models"
	"github.com/stretchr/testify/assert"
)

func TestCell_NeighbourAt(t *testing.T) {
	b := models.NewBoard()
	assert.Equal(t, "F6", b.CellAt("F5").NeighbourAt(models.NORTH).Position())
}

func TestCell_NeighbourAt_when_PositionIsInvalid_then_ReturnNil(t *testing.T) {
	b := models.NewBoard()
	assert.Nil(t, b.CellAt("H8").NeighbourAt(models.EAST))
}
