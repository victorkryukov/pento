package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(NewBoard(0, 1))
	assert.Nil(NewBoard(1, -1))
	assert.Equal(NewBoard(-1, -2).String(), "empty board")
	assert.Equal(NewBoard(2, 3).String(), "..\n..\n..\n")

	b := NewBoard(2, 3)
	f := Figure{Point{0, 0}, Point{1, 0}}
	assert.True(b.Place(f))
	assert.False(b.Place(f))
	assert.Equal(b.String(), "..\n..\naa\n")
	assert.True(b.Place(f.Move(Point{0, 1})))
	assert.Equal(b.String(), "..\nbb\naa\n")
	assert.False(b.Place(Figure{Point{0, 2}, Point{1, 2}, Point{0, 3}}))
	assert.Equal(b.String(), "..\nbb\naa\n")
	assert.True(b.PlaceAt(f, Point{0, 2}))
	assert.Equal(b.String(), "cc\nbb\naa\n")

	b = NewBoard(53, 1)
	for i := 0; i <= 51; i++ {
		assert.True(b.Place(Figure{Point{i, 0}}))
	}
	assert.Equal("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.\n", b.String())
	assert.True(b.Place(Figure{Point{52, 0}}))
	assert.Equal(b.String(), "too many figures on the board: 53")
	b.Unplace()
	b.Unplace()
	b.Unplace()
	assert.Equal("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWX...\n", b.String())
}
