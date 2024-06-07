package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	assert.NotNil(t, b)
	assert.Equal(t, 64, len(b.cells), "Board should have %d cells", 64)
	assert.Equal(t, "A1", b.cells["A1"].Position)
	assert.True(t, b.cells["A2"].neighbours[NORTH] != nil)
	assert.True(t, b.cells["A1"].neighbours[WEST] == nil)
	assert.True(t, b.cells["C3"].neighbours[SOUTH] == b.cells["C2"])
}
