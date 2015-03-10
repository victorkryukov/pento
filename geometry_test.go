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
