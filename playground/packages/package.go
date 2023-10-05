package packages

import "math"

//type (
//	X float64
//	Y float64
//	Z float64
//)

type Point struct {
	X float64
	Y float64
}

type ThreeDimensionPoint struct {
	Point
	Z float64
}
type Line struct {
	Start Point
	End   Point
}

type Circle struct {
	Point
	Radius float64
}

func (c *Circle) 面积() float64 {
	return math.Pi * c.Radius
}

func (l *Line) Length() float64 {

	return math.Hypot(l.End.X-l.Start.X, l.End.Y-l.Start.Y)
}

// Length  get the distance of between two point on an axis
func Length(s Point, e Point) (length float64) {
	return math.Hypot(s.X-e.X, s.Y-e.Y)
}
