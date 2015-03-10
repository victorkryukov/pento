package main

// Point is a point on a (x,y) plane with integer coordinates.
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

func main() {}
