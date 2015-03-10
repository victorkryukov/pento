package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Point{1, 0}.Rotate(), Point{0, -1})
	assert.Equal(Point{0, 1}.Rotate(), Point{1, 0})
	assert.Equal(Point{1, 1}.Mirror(), Point{1, 1})
	assert.Equal(Point{1, 0}.Mirror(), Point{0, 1})
	assert.Equal(Point{12, -5}.Mirror().Mirror(), Point{12, -5})
	assert.True(Point{12, -5}.Mirror().Mirror().Equal(Point{12, -5}))
}

func TestFigure(t *testing.T) {
	assert := assert.New(t)
	f := Figure{Point{1, 0}, Point{0, 1}}
	assert.True(f.Mirror().Equal(f))
	f1 := Figure{Point{1, 0}}
	assert.False(f1.Equal(f))
	assert.False(f.Equal(f1))
	assert.False(f1.Mirror().Equal(f1))
	assert.True(f1.Mirror().Mirror().Equal(f1))
	assert.False(f.Rotate().Equal(f))
	assert.False(f.Rotate().Rotate().Equal(f))
	assert.False(f.Rotate().Rotate().Rotate().Equal(f))
	assert.True(f.Rotate().Rotate().Rotate().Rotate().Equal(f))
	assert.Equal(f1.Move(Point{-1, -1}), f1.Rotate())
}

func TestFigureString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Figure{Point{1, 1}, Point{1, 2}, Point{1, 3}}.String(), "a\na\na\n")
	assert.Equal(Figure{Point{1, 3}, Point{1, 2}, Point{1, 1}}.String(), "a\na\na\n")
	assert.Equal(Figure{Point{1, 1}, Point{2, 1}, Point{1, 2}}.String(), "a.\naa\n")
	assert.Equal(Figure{Point{3, 2}, Point{2, 3}, Point{2, 2}}.String(), "a.\naa\n")
	assert.Equal(Figure{}.String(), "empty figure")
}

func TestFigureRecenter(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Figure{}.Recenter(), Figure{})
	f := Figure{Point{2, 3}, Point{1, 2}, Point{2, 2}}
	f1 := Figure{Point{0, 0}, Point{1, 1}, Point{1, 0}}
	assert.True(f.Recenter().Equal(f1), "\n%s\nrecenter should be equal to\n%s", f, f1)
	f = Figure{Point{1, 1}, Point{2, 2}, Point{1, 0}}
	f1 = Figure{Point{0, 0}, Point{0, 1}, Point{1, 2}}
	assert.True(f.Recenter().Equal(f1), "\n%s\nrecenter should be equal to\n%s", f, f1)
}
