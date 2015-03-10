package main

import (
	"bytes"
	"fmt"
)

// Board represents a rectangular board with some properly placed figures
type Board struct {
	SizeX, SizeY int
	Field        [][]bool // Tells whether (X,Y) coordinate is occupied
	Figures      []Figure // Figures placed so far
}

// NewBoard allocates memory for a new M x N board.
func NewBoard(m, n int) *Board {
	if m <= 0 || n <= 0 {
		return nil
	}
	b := &Board{SizeX: m, SizeY: n}
	b.Field = make([][]bool, m)
	for i := 0; i < m; i++ {
		b.Field[i] = make([]bool, n)
	}
	return b
}

// Place tries to place a figure on the board. It returns true and modifies
// the underlying board if placement was successful, otherwise it returns
// false and DOES NOT modify the board.
func (b *Board) Place(f Figure) bool {
	for _, p := range f {
		if p.X < 0 || p.X >= b.SizeX || p.Y < 0 || p.Y >= b.SizeY || b.Field[p.X][p.Y] {
			return false
		}
	}
	for _, p := range f {
		b.Field[p.X][p.Y] = true
	}
	b.Figures = append(b.Figures, f)
	return true
}

// String returns a board representation as a string. Empty cells
// are represented with '.', and figures each with it's own letter a-zA-Z.
func (b *Board) String() string {
	if b == nil || b.SizeX == 0 || b.SizeY == 0 {
		return "empty board"
	}
	if len(b.Figures) > 52 {
		return fmt.Sprintf("too many figures on the board: %d", len(b.Figures))
	}
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	field := make([][]byte, b.SizeX)
	for i := 0; i < b.SizeX; i++ {
		for j := 0; j < b.SizeY; j++ {
			field[i] = append(field[i], '.')
		}
	}
	for i, f := range b.Figures {
		for _, p := range f {
			field[p.X][p.Y] = letters[i]
		}
	}
	var buf bytes.Buffer
	for j := b.SizeY - 1; j >= 0; j-- {
		for i := 0; i < b.SizeX; i++ {
			buf.WriteByte(field[i][j])
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
