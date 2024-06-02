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
