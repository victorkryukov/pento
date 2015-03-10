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
	empty        int      // number of empty places on the board
}

// NewBoard allocates memory for a new M x N board.
func NewBoard(m, n int) *Board {
	if m <= 0 || n <= 0 {
		return nil
	}
	b := &Board{SizeX: m, SizeY: n}
	b.empty = m * n
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
	if len(f) > b.empty {
		return false
	}
	for _, p := range f {
		if p.X < 0 || p.X >= b.SizeX || p.Y < 0 || p.Y >= b.SizeY || b.Field[p.X][p.Y] {
			return false
		}
	}
	for _, p := range f {
		b.Field[p.X][p.Y] = true
	}
	b.empty -= len(f)
	b.Figures = append(b.Figures, f)
	return true
}

// PlaceAt is similar to Place, but places at point p as opposed to Point{0,0}.
func (b *Board) PlaceAt(f Figure, p Point) bool {
	return b.Place(f.Move(p))
}

// Unplace remove the figure that was last added from the board.
func (b *Board) Unplace() {
	if n := len(b.Figures); n > 0 {
		for _, p := range b.Figures[n-1] {
			b.Field[p.X][p.Y] = false
		}
		b.empty += len(b.Figures[n-1])
		b.Figures = b.Figures[:n-1]
	}
}

// Full returns true if board is full
func (b *Board) Full() bool {
	return b.empty == 0
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

var fillIterations int // Cumulative number of Fill calls
const printAfterNIterations = -1

// Fill tries to place each figure on the board and on success calls it recursively.
// When board is full, a handler is called.
func (b *Board) Fill(figures []Figure, handler func(*Board)) {
	if b.Full() {
		handler(b)
		return
	}
	for x := 0; x < b.SizeX; x++ {
		for y := 0; y < b.SizeY; y++ {
			if !b.Field[x][y] {
				//fmt.Printf("Trying to place a figure at (%d,%d) in\n%s", x+1, y+1, b)
				for _, f := range figures {
					if b.PlaceAt(f, Point{x, y}) {
						fillIterations++
						if printAfterNIterations > 0 && fillIterations%printAfterNIterations == 0 {
							fmt.Println(b)
						}
						b.Fill(figures, handler)
						b.Unplace()
					} else {
						//		fmt.Printf("doesn't fit:\n%s\n", f)
					}
				}
				return // Since we are trying to place pentos in the first available place
			}
		}
	}
}
