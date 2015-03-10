package main

// Point is a point/vector on a (x,y) plane with integer coordinates.
type Point struct {
	X, Y int
}

// Rotate rotates a point 90 degrees clock-wise.
func (p Point) Rotate() Point {
	return Point{p.Y, -p.X}
}

// Mirror returns point's reflection in x=y mirror.
func (p Point) Mirror() Point {
	return Point{p.Y, p.X}
}

// Equal returns true if two points are the same.
func (p Point) Equal(p1 Point) bool {
	return p.X == p1.X && p.Y == p1.Y
}

// Figure is a slice of points.
type Figure []Point

// Rotate rotates each point in a figure.
func (f Figure) Rotate() Figure {
	newf := Figure{}
	for _, p := range f {
		newf = append(newf, p.Rotate())
	}
	return newf
}

// Mirror mirrors each point in a figure.
func (f Figure) Mirror() Figure {
	newf := Figure{}
	for _, p := range f {
		newf = append(newf, p.Mirror())
	}
	return newf
}

// Move moves a figure by vector p
func (f Figure) Move(p Point) Figure {
	newf := Figure{}
	for _, pp := range f {
		newf = append(newf, Point{pp.X + p.X, pp.Y + p.Y})
	}
	return newf
}

// String returns a figure representation as a string, as if it
// was placed on a board starting in the bottom left corner.
func (f Figure) String() string {
	if len(f) == 0 {
		return "empty figure"
	}
	minX := f[0].X
	maxX := f[0].X
	minY := f[0].Y
	maxY := f[0].Y
	for i := 1; i < len(f); i++ {
		if minX > f[i].X {
			minX = f[i].X
		}
		if maxX < f[i].X {
			maxX = f[i].X
		}
		if minY > f[i].Y {
			minY = f[i].Y
		}
		if maxY < f[i].Y {
			maxY = f[i].Y
		}
	}
	b := NewBoard(maxX-minX+1, maxY-minY+1)
	b.PlaceAt(f, Point{-minX, -minY})
	return b.String()
}

// Equal returns true if two figures contain the same set of points,
// maybe in different order. It's very inefficient, but we only use it
// in tests.
func (f Figure) Equal(f1 Figure) bool {
	for _, p := range f {
		for _, p1 := range f1 {
			if p.Equal(p1) {
				goto next
			}
		}
		return false
	next:
	}
	for _, p1 := range f1 {
		for _, p := range f {
			if p.Equal(p1) {
				goto next1
			}
		}
		return false
	next1:
	}
	return true
}
