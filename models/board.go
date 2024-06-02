package models

import (
	"fmt"
	"strconv"
	"strings"
)

const dimension = 8

const (
	PAWN  int = 1
	KING  int = 2
	QUEEN int = 3
)

type board struct {
	cells [][]int
}

func NewBoard() *board {
	cells := make([][]int, dimension)

	for i := range cells {
		cells[i] = make([]int, dimension)
	}

	return &board{
		cells: cells,
	}
}

func (b *board) Put(p Piece, pos string) error {
	pos = strings.ToUpper(pos)

	if err := b.validate(pos); err != nil {
		return err
	}

	letter := rune(pos[0])
	num, _ := strconv.ParseInt(string(pos[1]), 10, 0)

	row, col := b.indexOf(letter, int(num))
	b.cells[row][col] = p.UID()
	return nil
}

func (b *board) validate(pos string) error {
	if len(pos) != 2 {
		return fmt.Errorf("invalid position format: %s", pos)
	}

	if !(pos[0] >= 'A' && pos[0] <= 'H') {
		return fmt.Errorf("invalid position first character: %s", pos)
	}

	num, err := strconv.ParseInt(string(pos[1]), 10, 0)
	if err != nil {
		return fmt.Errorf("invalid position second character type, expected an int: %s : %v", pos, err)
	}

	if !(num >= 1 && num <= 8) {
		return fmt.Errorf("invalid position second character: %s", pos)
	}

	return nil
}

func (b *board) indexOf(posLetter rune, posInt int) (int, int) {
	row := int(posLetter - 65)
	col := posInt - 1

	return row, col
}
